import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import { RecordingStatus } from '../../clients/api/api';
import Dropdown from '../../components/Dropdown/Dropdown';
import Error from '../../components/Error/Error';
import Loading from '../../components/Loading/Loading';
import RecordingListItem from '../../components/RecordingListItem/RecordingListItem';
import { useFetchRecordings } from '../../hooks/recording';
import styles from './RecordingsView.module.scss';

function RecordingsView() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const { recordings, error, loading, setStatus, status } = useFetchRecordings({
    status: 'upcoming',
    sort_key: 'startsAt',
  });

  if (loading) {
    return <Loading />;
  }

  if (error) {
    return <Error message={error} />;
  }

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
      />
    ));
  };

  return (
    <div className={styles.Recordings}>
      <div className={styles.header}>
        <Dropdown
          value={status}
          onChange={(value) => setStatus(value as RecordingStatus)}
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
      </div>
      <div className={styles.recordings}>{renderRecordings()}</div>
    </div>
  );
}

export default RecordingsView;
