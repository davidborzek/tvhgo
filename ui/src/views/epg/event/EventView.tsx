import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

import Error from '@/components/common/error/Error';
import EventChannelInfo from '@/components/epg/event/channelInfo/EventChannelInfo';
import EventInfo from '@/components/epg/event/info/EventInfo';
import EventRelated from '@/components/epg/event/related/EventRelated';
import { useFetchEvent } from '@/hooks/epg';
import { useManageRecordingByEvent } from '@/hooks/recording';

import styles from './EventView.module.scss';

function EventView() {
  const navigate = useNavigate();
  const params = useParams();
  const { fetch, error, event, relatedEvents } = useFetchEvent();
  const { createRecording, pending } = useManageRecordingByEvent();

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

    if (event.dvrUuid) {
      navigate(`/recordings/${event.dvrUuid}`);
      return;
    }

    await createRecording(event.id);
    fetchEvent();
  };

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
