import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useRevalidator,
  useSearchParams,
} from 'react-router-dom';

import { RecordingStatus, getRecordings } from '@/clients/api/api';
import Dropdown from '@/components/common/dropdown/Dropdown';
import { useManageRecordings } from '@/hooks/recording';
import Button from '@/components/common/button/Button';
import { c } from '@/utils/classNames';
import { ListResponse, Recording } from '@/clients/api/api.types';
import Checkbox from '@/components/common/checkbox/Checkbox';
import DeleteConfirmationModal from '@/components/common/deleteConfirmationModal/DeleteConfirmationModal';
import { usePagination } from '@/hooks/pagination';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import EmptyState from '@/components/common/emptyState/EmptyState';

import styles from './RecordingsView.module.scss';
import RecordingListItem from '@/components/recordings/listItem/RecordingListItem';
import { TestIds } from '@/__test__/ids';

const defaultLimit = 50;

export async function loader({ request }: LoaderFunctionArgs) {
  const query = new URL(request.url).searchParams;

  return getRecordings({
    status: (query.get('status') as RecordingStatus) || 'upcoming',
    sort_key: 'starts_at',
    limit: defaultLimit,
    offset: parseInt(query.get('offset')!!) || 0,
  });
}

export function Component() {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const revalidator = useRevalidator()
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

  const {entries, total} = useLoaderData() as ListResponse<Recording>;

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
          navigate(`/recordings/${recording.id}`);
        }}
        onSelection={(selected) => {
          setSelectedRecordings((prv) =>
            selected
              ? new Set([...prv, recording])
              : new Set([...prv].filter((v) => v.id != recording.id))
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
        revalidator.revalidate()
      });

      return;
    }

    removeRecordings([...selectedRecordings].map((rec) => rec.id)).then(() => {
      clearSelection();
      setConfirmationModalVisible(false);
       revalidator.revalidate()
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
          />

          <Checkbox
            onChange={(checked) =>
              checked
                ? setSelectedRecordings(new Set(entries))
                : clearSelection()
            }
            className={styles.selectAll}
            checked={
              entries.length > 0 &&
              selectedRecordings.size === entries.length
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
    </div>
  );
}

Component.displayName = 'RecordingsView';
