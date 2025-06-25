import { Channel, EpgEvent, ListResponse } from '@/clients/api/api.types';
import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useSearchParams,
} from 'react-router-dom';
import { getChannel, getEpgEvents } from '@/clients/api/api';

import ChannelViewer from '@/components/channels/viewer/ChannelViewer';
import EventChannelInfo from '@/components/epg/event/channelInfo/EventChannelInfo';
import GuideEvent from '@/components/epg/guide/event/GuideEvent';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import styles from './ChannelView.module.scss';
import { usePagination } from '@/hooks/pagination';
import { useState } from 'react';

const defaultLimit = 50;

export async function loader({ request, params }: LoaderFunctionArgs) {
  if (!params.id) {
    return;
  }

  const query = new URL(request.url).searchParams;

  const channel = getChannel(params.id);

  const events = getEpgEvents({
    channel: params.id,
    limit: defaultLimit,
    offset: parseInt(query.get('offset') || '0') || 0,
  });

  return Promise.all([channel, events]);
}

export const Component = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const [selectedEvent, setSelectedEvent] = useState<EpgEvent | null>(null);
  const [isViewerOpen, setIsViewerOpen] = useState(false);

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  const [channel, { entries, total }] = useLoaderData() as [
    Channel,
    ListResponse<EpgEvent>,
  ];

  const handleWatch = (event: EpgEvent) => {
    setSelectedEvent(event);
    setIsViewerOpen(true);
  };

  const handleCloseViewer = () => {
    setIsViewerOpen(false);
    setSelectedEvent(null);
  };

  return (
    <div className={styles.channel}>
      <div className={styles.header}>
        <EventChannelInfo
          channelName={channel.name}
          picon={`/api/picon/${channel.piconId}`}
        />
      </div>
      <div className={styles.events}>
        {entries.map((event) => (
          <GuideEvent
            key={event.id}
            eventId={event.id}
            title={event.title}
            subtitle={event.subtitle}
            description={event.description}
            endsAt={event.endsAt}
            startsAt={event.startsAt}
            onClick={(id) => {
              navigate(`/guide/events/${id}`);
            }}
            onWatch={() => handleWatch(event)}
            dvrState={event.dvrState}
            showProgress
            showDate
          />
        ))}
      </div>
      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(total)}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
      {selectedEvent && (
        <ChannelViewer
          event={selectedEvent}
          isOpen={isViewerOpen}
          onClose={handleCloseViewer}
        />
      )}
    </div>
  );
};

Component.displayName = 'ChannelView';
