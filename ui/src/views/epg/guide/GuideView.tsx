import { useNavigate, useSearchParams } from 'react-router-dom';
import { useEffect, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import moment from 'moment';

import { useFetchEpg } from '@/hooks/epg';
import { EpgChannel } from '@/clients/api/api.types';
import Error from '@/components/common/error/Error';
import { c } from '@/utils/classNames';
import EmptyState from '@/components/common/emptyState/EmptyState';

import styles from './GuideView.module.scss';
import GuideChannel from '@/components/epg/guide/channel/GuideChannel';
import GuideControls from '@/components/epg/guide/controls/GuideControls';
import GuideEventColumn from '@/components/epg/guide/eventColumn/GuideEventColumn';
import GuideNavigation from '@/components/epg/guide/navigation/GuideNavigation';

const parseStartDate = (start?: string | null) => {
  if (!start || start === 'today') {
    return;
  }

  return parseInt(start, 10);
};

const calculateEndDate = (end?: string | null) => {
  if (!end || end === 'today') {
    return moment().add(24, 'hour').unix();
  }

  return moment.unix(parseInt(end, 10)).endOf('day').unix();
};

function previousPage(oldOffset: number, limit: number, total: number): number {
  if (oldOffset >= limit) {
    return oldOffset - limit;
  }

  if (total % limit === 0) {
    return total - limit;
  }
  while (total % limit !== 0) {
    total--;
  }
  return total;
}

function nextPage(oldOffset: number, limit: number, total: number): number {
  if (oldOffset + limit < total) {
    return oldOffset + limit;
  }
  return 0;
}

function filterEpg(epg: EpgChannel[], search: string) {
  return epg.filter((e) =>
    e.channelName.toLowerCase().includes(search.toLocaleLowerCase())
  );
}

function GuideView() {
  const navigate = useNavigate();
  const containerRef = useRef<HTMLDivElement>(null);
  const [searchParams, setSearchParams] = useSearchParams();
  const { t } = useTranslation();

  const { events, setStartsAt, setEndsAt, error } = useFetchEpg({
    startsAt: parseStartDate(searchParams.get('day')),
    endsAt: calculateEndDate(searchParams.get('day')),
    sort_key: 'channelNumber',
    limit: 100,
  });

  const setDate = (start?: number, end?: number) => {
    setStartsAt(start);
    setEndsAt(end);
    containerRef.current?.scrollTo(0, 0);
  };

  const [limit, _setLimit] = useState(5);

  const limitRef = useRef(limit);
  const setLimit = (data: number) => {
    limitRef.current = data;
    _setLimit(data);
  };

  useEffect(() => {
    const limitResults = () => {
      const { current } = limitRef;
      const { innerWidth } = window;

      if (innerWidth > 1500 && current !== 5) {
        setLimit(5);
        return;
      }

      if (innerWidth > 1200 && innerWidth <= 1500 && current !== 4) {
        setLimit(4);
        return;
      }

      if (innerWidth > 800 && innerWidth <= 1200 && current !== 3) {
        setLimit(3);
        return;
      }
      if (innerWidth > 600 && innerWidth <= 800 && current !== 2) {
        setLimit(2);
        return;
      }

      if (innerWidth <= 600 && current !== 1) {
        setLimit(1);
        return;
      }
    };

    limitResults();
    window.addEventListener('resize', limitResults);

    return () => {
      window.removeEventListener('resize', limitResults);
    };
  }, []);

  useEffect(() => {
    const scrollPos = searchParams.get('pos');

    if (events && events.length > 0) {
      containerRef.current?.scrollTo(0, parseInt(scrollPos || '0'));
    }
  }, [events, searchParams]);

  const handleEventClick = (id: number) => {
    if (containerRef.current?.scrollTop !== undefined) {
      const pos = Math.floor(containerRef.current?.scrollTop);

      setSearchParams((prev) => {
        prev.set('pos', `${Math.floor(pos)}`);
        return prev;
      });
    }

    navigate(`/guide/events/${id}`, {
      preventScrollReset: true,
    });
  };

  const handleNextPageClick = () => {
    setSearchParams((prev) => {
      const offset = nextPage(
        parseInt(prev.get('offset') || '0'),
        limit,
        filteredEpg.length
      );

      prev.set('offset', `${offset}`);
      return prev;
    });
  };

  const handlePreviousPageClick = () => {
    setSearchParams((prev) => {
      const offset = previousPage(
        parseInt(prev.get('offset') || '0'),
        limit,
        filteredEpg.length
      );

      prev.set('offset', `${offset}`);
      return prev;
    });
  };

  const handleDayChange = (day: string) => {
    setSearchParams((prev) => {
      prev.set('day', day);
      return prev;
    });

    setDate(parseStartDate(day), calculateEndDate(day));
  };

  const handleSearch = (search: string) => {
    setSearchParams((prev) => {
      prev.set('search', search);
      prev.set('offset', '0');
      return prev;
    });

    containerRef.current?.scrollTo(0, 0);
  };

  const filteredEpg = filterEpg(events || [], searchParams.get('search') || '');

  const renderChannels = () => {
    const offset = parseInt(searchParams.get('offset') || '0');
    return filteredEpg
      .slice(offset, limit + offset)
      .map((event) => (
        <GuideChannel
          key={event.channelId}
          name={event.channelName}
          picon={`/api/picon/${event.piconId}`}
          number={event.channelNumber}
          onClick={() => navigate(`/channels/${event.channelId}`)}
        />
      ));
  };

  const renderEventColumns = () => {
    const offset = parseInt(searchParams.get('offset') || '0');
    return filteredEpg
      .slice(offset, limit + offset)
      .map((channel) => (
        <GuideEventColumn
          key={channel.channelId}
          events={channel.events}
          onClick={handleEventClick}
        />
      ));
  };

  if (error) {
    return <Error message={error} />;
  }

  const renderNavigation = () => {
    if (filteredEpg.length == 0) {
      return <></>;
    }

    return (
      <>
        <GuideNavigation type="left" onClick={handlePreviousPageClick} />
        <GuideNavigation type="right" onClick={handleNextPageClick} />
      </>
    );
  };

  return (
    <div className={styles.container} ref={containerRef}>
      <div className={styles.header}>
        <div className={styles.bar}>
          <GuideControls
            day={searchParams.get('day') || 'today'}
            search={searchParams.get('search') || ''}
            onDayChange={handleDayChange}
            onSearch={handleSearch}
          />
        </div>
        <div className={c(styles.segment, styles.channels)}>
          {renderChannels()}
          {renderNavigation()}
        </div>
      </div>

      {!events ? (
        <></>
      ) : events.length == 0 ? (
        <EmptyState title={t('no_epg')} />
      ) : (
        <div className={styles.segment}>{renderEventColumns()}</div>
      )}
    </div>
  );
}

export default GuideView;
