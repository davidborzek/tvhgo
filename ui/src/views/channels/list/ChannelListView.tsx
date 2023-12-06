import { useEffect, useRef } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { useFetchEvents } from '@/hooks/epg';
import ChannelListItem from '@/components/channels/listItem/ChannelListItem';
import Error from '@/components/common/error/Error';
import { usePagination } from '@/hooks/pagination';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import EmptyState from '@/components/common/emptyState/EmptyState';

import styles from './ChannelListView.module.scss';

const defaultLimit = 50;

function ChannelListView() {
  const ref = useRef<HTMLDivElement>(null);
  const { t } = useTranslation();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  const { events, total, error } = useFetchEvents({
    nowPlaying: true,
    limit,
    offset: getOffset(),
    sort_key: 'channelNumber',
    sort_dir: 'asc',
  });

  useEffect(() => {
    const scrollPos = searchParams.get('pos');

    if (ref.current?.scrollTo && events && events.length > 0) {
      ref.current?.scrollTo(0, parseInt(scrollPos || '0'));
    }
  }, [events, searchParams]);

  const renderChannels = () => {
    if (!events) {
      return <></>;
    }

    if (events && events.length === 0) {
      return (
        <EmptyState title={t('no_channels')} subtitle={t('no_channels_info')} />
      );
    }

    return events.map((event) => {
      return (
        <ChannelListItem
          key={event.channelId}
          event={event}
          onClick={(id) => {
            if (ref.current?.scrollTop !== undefined) {
              const pos = Math.floor(ref.current?.scrollTop);
              setSearchParams((prev) => {
                prev.set('pos', `${Math.floor(pos)}`);
                return prev;
              });
            }

            navigate(`/channels/${id}`);
          }}
        />
      );
    });
  };

  if (error) {
    return <Error message={error} />;
  }

  return (
    <div ref={ref} className={styles.container}>
      <div className={styles.channelList}>{renderChannels()}</div>
      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(total)}
        onPageChange={() => {
          setSearchParams((prev) => {
            prev.delete('pos');
            return prev;
          });
          ref.current?.scrollTo && ref.current?.scrollTo(0, 0);
        }}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
    </div>
  );
}

export default ChannelListView;
