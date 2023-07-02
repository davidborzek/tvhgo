import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { useFetchChannel } from '../../hooks/channel';
import { useEffect, useRef } from 'react';
import EventChannelInfo from '../../components/Event/EventChannelInfo/EventChannelInfo';

import styles from './ChannelView.module.scss';
import GuideEvent from '../../components/Guide/GuideEvent/GuideEvent';
import { usePagination } from '../../hooks/pagination';
import PaginationControls from '../../components/PaginationControls/PaginationControls';
import Error from '../../components/Error/Error';

const defaultLimit = 50;

const ChannelView = () => {
  const ref = useRef<HTMLDivElement>(null);
  const { id } = useParams();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  const { channel, events, total, error } = useFetchChannel(
    id,
    getOffset(),
    limit
  );

  useEffect(() => {
    const scrollPos = searchParams.get('pos');

    if (events.length > 0) {
      ref.current?.scrollTo(0, parseInt(scrollPos || '0'));
    }
  }, [events, searchParams]);

  if (error) {
    return <Error message={error} />;
  }

  if (!channel) {
    return <></>;
  }

  return (
    <div ref={ref} className={styles.channel}>
      <div className={styles.header}>
        <EventChannelInfo
          channelName={channel.name}
          picon={`/api/picon/${channel.piconId}`}
        />
      </div>
      <div className={styles.events}>
        {events.map((event) => (
          <GuideEvent
            eventId={event.id}
            title={event.title}
            subtitle={event.subtitle}
            description={event.description}
            endsAt={event.endsAt}
            startsAt={event.startsAt}
            onClick={(id) => {
              if (ref.current?.scrollTop !== undefined) {
                const pos = Math.floor(ref.current?.scrollTop);
                setSearchParams((prev) => {
                  prev.set('pos', `${Math.floor(pos)}`);
                  return prev;
                });
              }

              navigate(`/guide/events/${id}`, {
                preventScrollReset: true,
              });
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
        scrollTop={() => {
          setSearchParams((prev) => {
            prev.delete('pos');
            return prev;
          });
          ref.current?.scrollTo(0, 0);
        }}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
    </div>
  );
};

export default ChannelView;
