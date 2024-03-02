import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate, useSearchParams } from 'react-router-dom';

import { RecordingStatus } from '@/clients/api/api';
import Dropdown from '@/components/common/dropdown/Dropdown';
import Error from '@/components/common/error/Error';
import { useFetchRecordings, useManageRecordings } from '@/hooks/recording';
import Button from '@/components/common/button/Button';
import { c } from '@/utils/classNames';
import { Recording } from '@/clients/api/api.types';
import Checkbox from '@/components/common/checkbox/Checkbox';
import DeleteConfirmationModal from '@/components/common/deleteConfirmationModal/DeleteConfirmationModal';
import { usePagination } from '@/hooks/pagination';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import { useLoading } from '@/contexts/LoadingContext';
import EmptyState from '@/components/common/emptyState/EmptyState';

import styles from './RecordingsView.module.scss';
import RecordingListItem from '@/components/recordings/listItem/RecordingListItem';
import { TestIds } from '@/__test__/ids';
import ButtonLink from '@/components/common/button/ButtonLink';

const defaultLimit = 50;

function RecordingsView() {
  const ref = useRef<HTMLDivElement>(null);
  const { t } = useTranslation();
  const navigate = useNavigate();
  const { isLoading } = useLoading();
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

  const { recordings, error, fetch, total } = useFetchRecordings({
    status: getStatus(),
    sort_key: 'startsAt',
    limit,
    offset: getOffset(),
  });

  useEffect(() => {
    const scrollPos = queryParams.get('pos');

    if (ref.current?.scrollTo && recordings && recordings.length > 0) {
      ref.current?.scrollTo(0, parseInt(scrollPos || '0'));
    }
  }, [recordings, queryParams]);

  if (isLoading) {
    return <></>;
  }

  if (error) {
    return <Error message={error} />;
  }

  const getDeleteOrCancelButtonLabel = () => {
    return getStatus() === 'upcoming' ? t('cancel') : t('delete');
  };

  const getConfirmationModalTitle = () => {
    return getStatus() === 'upcoming'
      ? t('confirm_cancel_recordings')
      : t('confirm_delete_recordings');
  };

  const renderRecordings = () => {
    if (!recordings) {
      return <></>;
    }

    if (recordings.length === 0) {
      return <EmptyState title={t('no_recordings')} />;
    }

    return recordings.map((recording) => (
      <RecordingListItem
        key={recording.id}
        recording={recording}
        onClick={() => {
          if (ref.current?.scrollTop !== undefined) {
            const pos = Math.floor(ref.current?.scrollTop);
            setQueryParams((prev) => {
              prev.set('pos', `${Math.floor(pos)}`);
              return prev;
            });
          }

          navigate(`/recordings/${recording.id}`, { preventScrollReset: true });
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
        fetch();
      });

      return;
    }

    removeRecordings([...selectedRecordings].map((rec) => rec.id)).then(() => {
      clearSelection();
      setConfirmationModalVisible(false);
      fetch();
    });
  };

  return (
    <div ref={ref} className={styles.Recordings}>
      <DeleteConfirmationModal
        visible={confirmationModalVisible}
        onClose={() => setConfirmationModalVisible(false)}
        onConfirm={handleDeleteOrCancelRecordings}
        title={getConfirmationModalTitle()}
        buttonTitle={getDeleteOrCancelButtonLabel()}
        pending={pending}
      />

      <div className={styles.header}>
        <div className={styles.leftActions}>
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

          <Button label={t('add')} onClick={() => navigate('create')} />
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
          />

          <Checkbox
            onChange={(checked) =>
              checked
                ? setSelectedRecordings(new Set(recordings))
                : clearSelection()
            }
            className={styles.selectAll}
            checked={
              !!recordings &&
              recordings.length > 0 &&
              selectedRecordings.size === recordings.length
            }
            indeterminate={selectedRecordings.size > 0}
            disabled={!!recordings && recordings.length < 1}
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
          setQueryParams((prev) => {
            prev.delete('pos');
            return prev;
          });
          ref.current?.scrollTo(0, 0);
          clearSelection();
        }}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
    </div>
  );
}

export default RecordingsView;
