import { ListResponse, UserResponse } from '@/clients/api/api.types';
import {
  LoaderFunctionArgs,
  Outlet,
  useLoaderData,
  useLocation,
  useNavigate,
  useRevalidator,
  useSearchParams,
} from 'react-router-dom';
import { useEffect, useState } from 'react';

import Badge from '@/components/common/badge/Badge';
import Button from '@/components/common/button/Button';
import DeleteConfirmationModal from '@/components/common/deleteConfirmationModal/DeleteConfirmationModal';
import EmptyState from '@/components/common/emptyState/EmptyState';
import Headline from '@/components/common/headline/Headline';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import Table from '@/components/common/table/Table';
import TableBody from '@/components/common/table/TableBody';
import TableCell from '@/components/common/table/TableCell';
import TableHead from '@/components/common/table/TableHead';
import TableHeadCell from '@/components/common/table/TableHeadCell';
import TableRow from '@/components/common/table/TableRow';
import { UserListRefreshStates } from './states';
import { getUsers } from '@/clients/api/api';
import styles from './UserListView.module.scss';
import { t } from 'i18next';
import { useDeleteUser } from '@/hooks/user';
import { usePagination } from '@/hooks/pagination';

const defaultLimit = 50;

export async function loader({ request }: LoaderFunctionArgs) {
  const query = new URL(request.url).searchParams;

  const users = getUsers({
    limit: defaultLimit,
    offset: parseInt(query.get('offset')!) || 0,
  });

  return Promise.all([users]);
}

export const Component = () => {
  const navigate = useNavigate();
  const { state } = useLocation();
  const { revalidate } = useRevalidator();

  const [userToDelete, setUserToDelete] = useState<number | null>(null);

  const [deleting, deleteUser] = useDeleteUser();

  const [queryParams, setQueryParams] = useSearchParams();
  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, queryParams, setQueryParams);
  const [users] = useLoaderData() as [ListResponse<UserResponse>];

  useEffect(() => {
    switch (state) {
      case UserListRefreshStates.CREATED:
        revalidate();
        break;
    }
  }, [state, revalidate]);

  const renderUsers = () => {
    // This should normally never happen, but just in case.
    if (users.entries.length === 0) {
      return <EmptyState title={t('no_users_found')} />;
    }

    return (
      <Table>
        <TableHead>
          <TableHeadCell>{t('username')}</TableHeadCell>
          <TableHeadCell className={styles.name}>{t('name')}</TableHeadCell>
          <TableHeadCell className={styles.email}>{t('email')}</TableHeadCell>
          <TableHeadCell className={styles.twofa}>{t('2fa')}</TableHeadCell>
          <TableHeadCell className={styles.created}>
            {t('created')}
          </TableHeadCell>
          <TableHeadCell />
        </TableHead>
        <TableBody>
          {users.entries.map((user) => (
            <TableRow key={user.id}>
              <TableCell>{user.username}</TableCell>
              <TableCell className={styles.name}>{user.displayName}</TableCell>
              <TableCell className={styles.email}>{user.email}</TableCell>
              <TableCell className={styles.twofa}>
                {user.twoFactor ? (
                  <Badge color="default">{t('yes')}</Badge>
                ) : (
                  <Badge color="failure">{t('no')}</Badge>
                )}
              </TableCell>
              <TableCell className={styles.created}>
                {t('date', { date: user.createdAt })}
              </TableCell>
              <TableCell className={styles.delete}>
                <Button
                  style="red"
                  size="small"
                  label={t('delete')}
                  quiet
                  onClick={() => setUserToDelete(user.id)}
                />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    );
  };

  return (
    <div className={styles.users}>
      <DeleteConfirmationModal
        visible={userToDelete !== null}
        onClose={() => setUserToDelete(null)}
        onConfirm={() => {
          if (userToDelete !== null) {
            deleteUser(userToDelete);
            setUserToDelete(null);
          }
        }}
        title={t('confirm_delete_user')}
        buttonTitle={t('delete')}
        pending={deleting}
      />

      <div className={styles.header}>
        <Headline>{t('users')}</Headline>
        <Button
          label={t('create')}
          quiet
          onClick={() => navigate('/settings/users/create')}
        />
      </div>
      {renderUsers()}
      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(users.total)}
        limit={limit}
        offset={getOffset()}
        total={users.total}
      />
      <Outlet />
    </div>
  );
};

Component.displayName = 'UserListView';
