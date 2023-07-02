import { useEffect, useState } from 'react';
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

function RecordingsView() {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const [queryParams, setQueryParams] = useSearchParams();
  const [selectedRecordings, setSelectedRecordings] = useState<Set<Recording>>(
    new Set()
  );
  const clearSelection = () => setSelectedRecordings(new Set());

  const { stopAndCancelRecordings, removeRecordings, pending } =
    useManageRecordings();

  const { recordings, error, setStatus, status, fetch } = useFetchRecordings({
    status: 'upcoming',
    sort_key: 'startsAt',
  });

  useEffect(() => {
    setStatus((queryParams.get('status') as RecordingStatus) || 'upcoming');
  }, [queryParams]);

  if (error) {
    return <Error message={error} />;
  }

  const getDeleteOrCancelButtonLabel = () => {
    return status === 'upcoming' ? t('cancel') : t('delete');
  };

  const renderRecordings = () => {
    if (recordings.length === 0) {
      return <div className={styles.emptyState}>{t('no_recordings')}</div>;
    }

    return recordings.map((recording) => (
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
    if (status === 'upcoming') {
      const stopIds = [...selectedRecordings]
        .filter((rec) => rec.status === 'recording')
        .map((rec) => rec.id);

      const cancelIds = [...selectedRecordings]
        .filter((rec) => rec.status !== 'recording')
        .map((rec) => rec.id);

      stopAndCancelRecordings(stopIds, cancelIds, () => {
        clearSelection();
        fetch();
      });

      return;
    }

    removeRecordings(
      [...selectedRecordings].map((rec) => rec.id),
      () => {
        clearSelection();
        fetch();
      }
    );
  };

  return (
    <div className={styles.Recordings}>
      <div className={styles.header}>
        <Dropdown
          value={status}
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
            onClick={handleDeleteOrCancelRecordings}
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
              recordings.length > 0 &&
              selectedRecordings.size === recordings.length
            }
            indeterminate={selectedRecordings.size > 0}
            disabled={recordings.length < 1}
          />
        </div>
      </div>
      <div className={styles.recordings}>{renderRecordings()}</div>
    </div>
  );
}

export default RecordingsView;
