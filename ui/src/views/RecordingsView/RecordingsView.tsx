import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { RecordingStatus } from '../../clients/api/api';
import Dropdown from '../../components/Dropdown/Dropdown';
import Error from '../../components/Error/Error';
import RecordingListItem from '../../components/RecordingListItem/RecordingListItem';
import { useFetchRecordings, useManageRecordings } from '../../hooks/recording';
import styles from './RecordingsView.module.scss';
import Button from '../../components/Button/Button';
import { c } from '../../utils/classNames';
import { Recording } from '../../clients/api/api.types';
import Checkbox from '../../components/Checkbox/Checkbox';
import DeleteConfirmationModal from '../../components/DeleteConfirmationModal/DeleteConfirmationModal';
import { usePagination } from '../../hooks/pagination';
import PaginationControls from '../../components/PaginationControls/PaginationControls';
import { useLoading } from '../../contexts/LoadingContext';
import EmptyState from '../../components/EmptyState/EmptyState';

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

    if (recordings && recordings.length > 0) {
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

      stopAndCancelRecordings(stopIds, cancelIds, () => {
        clearSelection();
        setConfirmationModalVisible(false);
        fetch();
      });

      return;
    }

    removeRecordings(
      [...selectedRecordings].map((rec) => rec.id),
      () => {
        clearSelection();
        setConfirmationModalVisible(false);
        fetch();
      }
    );
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
        <Dropdown
          value={getStatus()}
          onChange={(value) => {
            clearSelection();
            setQueryParams({
              status: value,
            });
          }}
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
