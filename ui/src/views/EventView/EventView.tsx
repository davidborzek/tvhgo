import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Error from '../../components/Error/Error';
import EventChannelInfo from '../../components/Event/EventChannelInfo/EventChannelInfo';
import EventInfo from '../../components/Event/EventInfo/EventInfo';
import EventRelated from '../../components/Event/EventRelated/EventRelated';
import Loading from '../../components/Loading/Loading';
import { useFetchEvent } from '../../hooks/epg';
import { useManageRecordingByEvent } from '../../hooks/recording';
import styles from './EventView.module.scss';

function EventView() {
  const params = useParams();
  const { fetch, error, event, relatedEvents, loading } = useFetchEvent();
  const { createRecording, stopRecording, cancelRecording, pending } =
    useManageRecordingByEvent();

  const fetchEvent = () => {
    const id = params['id'];
    if (id) {
      fetch(parseInt(id));
    }
  };

  useEffect(() => {
    fetchEvent();
  }, [params]);

  const handleOnRecord = async () => {
    if (!event) {
      return;
    }

    if (event.dvrState === 'scheduled') {
      await cancelRecording(event.dvrUuid!);
    } else if (event.dvrState === 'recording') {
      await stopRecording(event.dvrUuid!);
    } else {
      await createRecording(event.id);
    }

    fetchEvent();
  };

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
      <EventInfo
        event={event}
        handleOnRecord={handleOnRecord}
        pending={pending}
      />
      <EventRelated relatedEvents={relatedEvents} />
    </div>
  );
}

export default EventView;
