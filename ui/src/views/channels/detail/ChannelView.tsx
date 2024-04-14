import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useParams,
  useSearchParams,
} from 'react-router-dom';
import { useRef } from 'react';

import EventChannelInfo from '@/components/epg/event/channelInfo/EventChannelInfo';
import { usePagination } from '@/hooks/pagination';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';

import styles from './ChannelView.module.scss';
import GuideEvent from '@/components/epg/guide/event/GuideEvent';
import { getChannel, getEpgEvents } from '@/clients/api/api';
import { Channel, EpgEvent, ListResponse } from '@/clients/api/api.types';

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
  const ref = useRef<HTMLDivElement>(null);
  const { id } = useParams();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  const [channel, { entries, total }] = useLoaderData() as [
    Channel,
    ListResponse<EpgEvent>,
  ];

  return (
    <div ref={ref} className={styles.channel}>
      <div className={styles.header}>
        <EventChannelInfo
          channelName={channel.name}
          picon={`/api/picon/${channel.piconId}`}
        />
      </div>
      <div className={styles.events}>
        {entries.map((event) => (
          <GuideEvent
            eventId={event.id}
            title={event.title}
            subtitle={event.subtitle}
            description={event.description}
            endsAt={event.endsAt}
            startsAt={event.startsAt}
            onClick={(id) => {
              navigate(`/guide/events/${id}`);
            }}
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
    </div>
  );
};

Component.displayName = 'ChannelView';
