import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Error from '../../components/Error/Error';
import EventChannelInfo from '../../components/Event/EventChannelInfo/EventChannelInfo';
import EventInfo from '../../components/Event/EventInfo/EventInfo';
import EventRelated from '../../components/Event/EventRelated/EventRelated';
import Loading from '../../components/Loading/Loading';
import { useFetchEvent } from '../../hooks/epg';
import styles from './Event.module.scss';

function Event() {
  const params = useParams();
  const { fetch, error, event, relatedEvents, loading } = useFetchEvent();

  useEffect(() => {
    const id = params['id'];
    if (id) {
      fetch(parseInt(id));
    }
  }, [params]);

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
      <EventChannelInfo
        channelName={event.channelName}
        picon={`/api/picon/${event.piconId}`}
      />
      <EventInfo event={event} />
      <EventRelated relatedEvents={relatedEvents} />
    </div>
  );
}

export default Event;
