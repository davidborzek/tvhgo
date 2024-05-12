import { ListResponse, Recording } from '@/clients/api/api.types';
import {
  LoaderFunctionArgs,
  Outlet,
  useLoaderData,
  useLocation,
  useNavigate,
  useRevalidator,
  useSearchParams,
} from 'react-router-dom';
import { RecordingStatus, getRecordings } from '@/clients/api/api';
import { useEffect, useMemo, useState } from 'react';

import Button from '@/components/common/button/Button';
import Checkbox from '@/components/common/checkbox/Checkbox';
import DeleteConfirmationModal from '@/components/common/deleteConfirmationModal/DeleteConfirmationModal';
import Dropdown from '@/components/common/dropdown/Dropdown';
import EmptyState from '@/components/common/emptyState/EmptyState';
import { LargeArrowLeftIcon } from '@/assets';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import RecordingListItem from '@/components/recordings/listItem/RecordingListItem';
import { TestIds } from '@/__test__/ids';
import { c } from '@/utils/classNames';
import styles from './RecordingsView.module.scss';
import { useManageRecordings } from '@/hooks/recording';
import { usePagination } from '@/hooks/pagination';
import { useTranslation } from 'react-i18next';

export enum RecordingsViewRefreshStates {
  CREATED = 'recording_created',
}

const defaultLimit = 50;

export async function loader({ request }: LoaderFunctionArgs) {
  const query = new URL(request.url).searchParams;

  return getRecordings({
    limit: defaultLimit,
    offset: parseInt(query.get('offset')!) || 0,
    // eslint-disable-next-line camelcase
    sort_key: query.get('sortKey') || 'startsAt',
    // eslint-disable-next-line camelcase
    sort_dir: query.get('sortDir') || 'asc',
    status: (query.get('status') as RecordingStatus) || 'upcoming',
  });
}

export function Component() {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const { state } = useLocation();
  const { revalidate } = useRevalidator();
  const [queryParams, setQueryParams] = useSearchParams();
  const [selectedRecordings, setSelectedRecordings] = useState<Set<Recording>>(
    new Set()
  );
  const clearSelection = () => setSelectedRecordings(new Set());
  const [confirmationModalVisible, setConfirmationModalVisible] =
    useState<boolean>(false);

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, queryParams, setQueryParams);

  const { stopAndCancelRecordings, removeRecordings, pending } =
    useManageRecordings();

  const getStatus = () =>
    (queryParams.get('status') as RecordingStatus) || 'upcoming';

  const sortKey = useMemo(
    () => queryParams.get('sortKey') || 'startsAt',
    [queryParams]
  );

  const sortDir = useMemo(
    () => queryParams.get('sortDir') || 'asc',
    [queryParams]
  );

  const { entries, total } = useLoaderData() as ListResponse<Recording>;

  useEffect(() => {
    switch (state) {
      case RecordingsViewRefreshStates.CREATED:
        revalidate();
        break;
    }
  }, [state, revalidate]);

  const getDeleteOrCancelButtonLabel = () => {
    return getStatus() === 'upcoming' ? t('cancel') : t('delete');
  };

  const getConfirmationModalTitle = () => {
    return getStatus() === 'upcoming'
      ? t('confirm_cancel_recordings')
      : t('confirm_delete_recordings');
  };

  const renderRecordings = () => {
    if (entries.length === 0) {
      return <EmptyState title={t('no_recordings')} />;
    }

    return entries.map((recording) => (
      <RecordingListItem
        key={recording.id}
        recording={recording}
        onClick={() => {
          navigate(`/dvr/recordings/${recording.id}`);
        }}
        onSelection={(selected) => {
          setSelectedRecordings((prv) =>
            selected
              ? new Set([...prv, recording])
              : new Set([...prv].filter((v) => v.id !== recording.id))
          );
        }}
        selected={selectedRecordings.has(recording)}
      />
    ));
  };

  const handleDeleteOrCancelRecordings = () => {
    if (getStatus() === 'upcoming') {
      const stopIds = [...selectedRecordings]
        .filter((rec) => rec.status === 'recording')
        .map((rec) => rec.id);

      const cancelIds = [...selectedRecordings]
        .filter((rec) => rec.status !== 'recording')
        .map((rec) => rec.id);

      stopAndCancelRecordings(stopIds, cancelIds).then(() => {
        clearSelection();
        setConfirmationModalVisible(false);
        revalidate();
      });

      return;
    }

    removeRecordings([...selectedRecordings].map((rec) => rec.id)).then(() => {
      clearSelection();
      setConfirmationModalVisible(false);
      revalidate();
    });
  };

  return (
    <div className={styles.Recordings}>
      <DeleteConfirmationModal
        visible={confirmationModalVisible}
        onClose={() => setConfirmationModalVisible(false)}
        onConfirm={handleDeleteOrCancelRecordings}
        title={getConfirmationModalTitle()}
        buttonTitle={getDeleteOrCancelButtonLabel()}
        pending={pending}
      />

      <div className={styles.header}>
        <div className={styles.headerLeft}>
          <Dropdown
            value={getStatus()}
            onChange={(value) => {
              clearSelection();
              setQueryParams({
                status: value,
              });
            }}
            testID={TestIds.RECORDINGS_STATUS_DROPDOWN}
            options={[
              {
                title: t('upcoming'),
                value: 'upcoming',
              },
              {
                title: t('finished'),
                value: 'finished',
              },
              {
                title: t('failed'),
                value: 'failed',
              },
              {
                title: t('removed'),
                value: 'removed',
              },
            ]}
          />

          <div className={styles.sort}>
            <Dropdown
              value={sortKey}
              onChange={(value) => {
                setQueryParams((prev) => {
                  prev.set('sortKey', value);
                  return prev;
                });
              }}
              testID={TestIds.RECORDINGS_SORT_DROPDOWN}
              options={[
                {
                  title: t('sort_by.start_date'),
                  value: 'startsAt',
                },
                {
                  title: t('sort_by.end_date'),
                  value: 'endsAt',
                },
                {
                  title: t('sort_by.title'),
                  value: 'title',
                },
                {
                  title: t('sort_by.channel'),
                  value: 'channelName',
                },
              ]}
            />

            <Button
              quiet
              style="text"
              icon={
                <LargeArrowLeftIcon
                  className={c(
                    styles.sortDirIcon,
                    sortDir === 'desc' ? styles.desc : ''
                  )}
                />
              }
              onClick={() => {
                setQueryParams((prev) => {
                  prev.set(
                    'sortDir',
                    prev.get('sortDir') === 'desc' ? 'asc' : 'desc'
                  );
                  return prev;
                });
              }}
              testID={TestIds.RECORDINGS_SORT_DIR_BUTTON}
            />
          </div>

          <Button
            label={t('create_recording')}
            style="blue"
            className={c()}
            quiet
            onClick={() => navigate('/dvr/recordings/create')}
          />
        </div>

        <div className={styles.actions}>
          <Button
            label={getDeleteOrCancelButtonLabel()}
            onClick={() => setConfirmationModalVisible(true)}
            style="red"
            disabled={pending}
            className={c(
              styles.deleteButton,
              selectedRecordings.size > 0 ? styles.deleteButtonVisible : ''
            )}
            testID={TestIds.DELETE_CANCEL_RECORDINGS_BUTTON}
            quiet
          />
          <Checkbox
            onChange={(checked) =>
              checked
                ? setSelectedRecordings(new Set(entries))
                : clearSelection()
            }
            className={styles.selectAll}
            checked={
              entries.length > 0 && selectedRecordings.size === entries.length
            }
            indeterminate={selectedRecordings.size > 0}
            disabled={entries.length < 1}
            testId={TestIds.SELECT_ALL_RECORDINGS_CHECKBOX}
          />
        </div>
      </div>
      <div className={styles.recordings}>{renderRecordings()}</div>
      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(total)}
        onPageChange={() => {
          clearSelection();
        }}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
      <Outlet />
    </div>
  );
}

Component.displayName = 'RecordingsView';
