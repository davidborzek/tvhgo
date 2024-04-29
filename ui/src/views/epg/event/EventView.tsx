import { DVRConfig, EpgEvent } from '@/clients/api/api.types';
import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useRevalidator,
} from 'react-router-dom';
import {
  getDVRConfigs,
  getEpgEvent,
  getRelatedEpgEvents,
} from '@/clients/api/api';

import DVRProfileSelectModal from '@/modals/dvr/profile/DVRProfileSelectModal';
import EventChannelInfo from '@/components/epg/event/channelInfo/EventChannelInfo';
import EventInfo from '@/components/epg/event/info/EventInfo';
import EventRelated from '@/components/epg/event/related/EventRelated';
import styles from './EventView.module.scss';
import { useManageRecordingByEvent } from '@/hooks/recording';
import { useState } from 'react';

export async function loader({ params }: LoaderFunctionArgs) {
  if (!params.id) {
    return;
  }

  const id = parseInt(params.id);
  if (!id) {
    return;
  }

  const [event, related, dvrProfiles] = await Promise.all([
    getEpgEvent(id),
    getRelatedEpgEvents(id),
    getDVRConfigs(),
  ]);

  return [event, related.entries.filter((r) => r.id !== id), dvrProfiles];
}

export function Component() {
  const navigate = useNavigate();
  const revalidator = useRevalidator();
  const { createRecording, pending } = useManageRecordingByEvent();
  const [event, relatedEvents, dvrProfiles] = useLoaderData() as [
    EpgEvent,
    Array<EpgEvent>,
    Array<DVRConfig>,
  ];

  const [modalVisible, setModalVisible] = useState(false);

  const handleCreateRecording = async (profileId?: string) => {
    await createRecording(event.id, profileId);
    revalidator.revalidate();
  };

  const handleOnRecord = async () => {
    if (!event) {
      return;
    }

    if (event.dvrUuid) {
      navigate(`/dvr/recordings/${event.dvrUuid}`);
      return;
    }

    if (dvrProfiles.length > 1) {
      setModalVisible(true);
      return;
    }

    handleCreateRecording();
  };

  return (
    <div className={styles.Event}>
      <DVRProfileSelectModal
        disableBackdropClose
        disableEscapeClose
        visible={modalVisible}
        onClose={() => setModalVisible(false)}
        profiles={dvrProfiles}
        maxWidth="25rem"
        handleCreateRecording={handleCreateRecording}
      />

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
