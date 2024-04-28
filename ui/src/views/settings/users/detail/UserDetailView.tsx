import { LoaderFunctionArgs, useLoaderData } from 'react-router-dom';
import { Session, UserResponse } from '@/clients/api/api.types';
import { getSessions, getUser } from '@/clients/api/api';

import Badge from '@/components/common/badge/Badge';
import Button from '@/components/common/button/Button';
import DeleteConfirmationModal from '@/components/common/deleteConfirmationModal/DeleteConfirmationModal';
import Headline from '@/components/common/headline/Headline';
import Pair from '@/components/common/pairList/Pair/Pair';
import PairKey from '@/components/common/pairList/PairKey/PairKey';
import PairList from '@/components/common/pairList/PairList';
import PairValue from '@/components/common/pairList/PairValue/PairValue';
import SessionList from '@/components/settings/sessionList/SessionList';
import styles from './UserDetailView.module.scss';
import { useDeleteUser } from '@/hooks/user';
import { useManageSessions } from '@/hooks/session';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export async function loader({ params }: LoaderFunctionArgs) {
  if (!params.id) {
    return [];
  }

  const id = parseInt(params.id);
  const user = getUser(id);

  const sessions = getSessions(id);

  return Promise.all([user, sessions]);
}

export const Component = () => {
  const { t } = useTranslation();
  const [user, sessions] = useLoaderData() as [UserResponse, Session[]];
  const [deleteConfirm, setDeleteConfirm] = useState(false);
  const [deleting, deleteUser] = useDeleteUser();
  const { revokeUserSession } = useManageSessions();

  return (
    <div className={styles.view}>
      <div className={styles.actions}>
        <Button
          onClick={() => setDeleteConfirm(true)}
          style="red"
          size="small"
          label={t('delete')}
          quiet
        />
      </div>

      <DeleteConfirmationModal
        visible={deleteConfirm}
        onClose={() => setDeleteConfirm(false)}
        onConfirm={() => {
          deleteUser(user.id);
          setDeleteConfirm(false);
        }}
        title={t('confirm_delete_user')}
        buttonTitle={t('delete')}
        pending={deleting}
      />

      <Headline>
        {t('user')}: {user.username}
      </Headline>
      <PairList>
        <Pair>
          <PairKey>{t('name')}</PairKey>
          <PairValue>{user.username}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('email')}</PairKey>
          <PairValue>{user.email}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('created')}</PairKey>
          <PairValue>{t('date', { date: user.createdAt })}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('2fa')}</PairKey>
          <PairValue>
            {user.twoFactor ? (
              <Badge color="default">{t('yes')}</Badge>
            ) : (
              <Badge color="failure">{t('no')}</Badge>
            )}
          </PairValue>
        </Pair>
      </PairList>

      <SessionList
        sessions={sessions}
        onRevoke={(id) => revokeUserSession(user.id, id)}
      />
    </div>
  );
};

Component.displayName = 'UserDetailView';
