import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useRevalidator,
} from 'react-router-dom';

import EventChannelInfo from '@/components/epg/event/channelInfo/EventChannelInfo';
import EventInfo from '@/components/epg/event/info/EventInfo';
import EventRelated from '@/components/epg/event/related/EventRelated';
import { useManageRecordingByEvent } from '@/hooks/recording';

import styles from './EventView.module.scss';
import { getEpgEvent, getRelatedEpgEvents } from '@/clients/api/api';
import { EpgEvent } from '@/clients/api/api.types';

export async function loader({ params }: LoaderFunctionArgs) {
  if (!params.id) {
    return;
  }

  const id = parseInt(params.id);
  if (!id) {
    return;
  }

  const [event, related] = await Promise.all([
    getEpgEvent(id),
    getRelatedEpgEvents(id),
  ]);

  return [event, related.entries.filter((r) => r.id !== id)];
}

export function Component() {
  const navigate = useNavigate();
  const revalidator = useRevalidator();
  const { createRecording, pending } = useManageRecordingByEvent();
  const [event, relatedEvents] = useLoaderData() as [EpgEvent, Array<EpgEvent>];

  const handleOnRecord = async () => {
    if (!event) {
      return;
    }

    if (event.dvrUuid) {
      navigate(`/recordings/${event.dvrUuid}`);
      return;
    }

    await createRecording(event.id);
    revalidator.revalidate();
  };

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

Component.displayName = 'EventView';
