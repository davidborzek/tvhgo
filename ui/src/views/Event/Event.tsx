import { useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { Navigate, useParams } from 'react-router-dom';
import Error from '../../components/Error/Error';
import Image from '../../components/Image/Image';
import Loading from '../../components/Loading/Loading';
import { useFetchEvent } from '../../hooks/epg';
import { parseTime } from '../../utils/time';
import styles from './Event.module.scss';

function Event() {
  const { t } = useTranslation();

  const params = useParams();
  const { fetch, error, event, loading } = useFetchEvent();

  useEffect(() => {
    const id = params['id'];
    if (id) {
      fetch(parseInt(id));
    }
  }, []);

  function renderDatetime(startsAt: number, endsAt: number): string {
    const date = new Date(startsAt * 1000);
    const day = date.getDay();

    const minutesLeft = Math.floor((endsAt - startsAt) / 60);
    const start = parseTime(startsAt);

    return `${t(
      `weekday_${day}`
    )}, ${date.toLocaleDateString()} • ${start} • ${minutesLeft} Min.`;
  }

  if (loading) {
    return <Loading />;
  }

  if (error) {
    return <Error message={error} />;
  }

  if (!event) {
    return <></>;
  }

  return (
    <div className={styles.Event}>
      <div className={styles.channelInfo}>
        <span className={styles.channelName}>{event.channelName}</span>
        <Image
          title={event.channelName}
          className={styles.picon}
          src={`/api/picon/${event.piconId}`}
        />
      </div>
      <h1>{event.title}</h1>
      <span>{renderDatetime(event.startsAt, event.endsAt)}</span>
      <p>{event.subtitle}</p>
      <span>{event.description}</span>
    </div>
  );
}

export default Event;
