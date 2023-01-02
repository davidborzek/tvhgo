import { useFetchChannelEvents } from '../../hooks/epg';
import styles from './GuideView.module.scss';
import moment from 'moment';
import { useEffect, useRef, useState } from 'react';
import GuideChannel from '../../components/Guide/GuideChannel/GuideChannel';
import GuideEventColumn from '../../components/Guide/GuideEventColumn/GuideEventColumn';
import GuideNavigation from '../../components/Guide/GuideNavigation/GuideNavigation';
import { EpgChannel } from '../../clients/api/api.types';
import GuideControls from '../../components/Guide/GuideControls/GuideControls';
import Error from '../../components/Error/Error';
import { useNavigate } from 'react-router-dom';
import { c } from '../../utils/classNames';

const SCROLL_PERSIST_KEY = 'tvhgo_guide_scroll_position';

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

  const { events, setStartsAt, setEndsAt, loading, error } =
    useFetchChannelEvents({
      endsAt: moment().add(24, 'hour').unix(),
      sort_key: 'channelNumber',
      limit: 100,
    });

  const setDate = (start?: number, end?: number) => {
    setStartsAt(start);
    setEndsAt(end);
    containerRef.current?.scrollTo(0, 0);
  };

  const [search, setSearch] = useState('');

  const [offset, setOffset] = useState(0);
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
    const scrollPos = localStorage.getItem(SCROLL_PERSIST_KEY);
    if (scrollPos && events.length > 0) {
      containerRef.current?.scrollTo(0, parseInt(scrollPos));
      localStorage.removeItem(SCROLL_PERSIST_KEY);
    }
  }, [events]);

  const filteredEpg = filterEpg(events, search);

  const handleDayChange = (dateString: string) => {
    if (dateString === 'today') {
      setDate(undefined, moment().endOf('day').unix());
      return;
    }

    const date = moment.unix(parseInt(dateString, 10));
    setDate(date.unix(), date.endOf('day').unix());
  };

  const renderChannels = () => {
    return filteredEpg
      .slice(offset, limit + offset)
      .map((event) => (
        <GuideChannel
          key={event.channelId}
          name={event.channelName}
          picon={`/api/picon/${event.piconId}`}
          number={event.channelNumber}
        />
      ));
  };

  const renderEventColumns = () => {
    return filteredEpg.slice(offset, limit + offset).map((channel) => (
      <GuideEventColumn
        key={channel.channelId}
        events={channel.events}
        onClick={(id) => {
          localStorage.setItem(
            SCROLL_PERSIST_KEY,
            `${containerRef.current?.scrollTop || ''}`
          );
          navigate(`/guide/events/${id}`, {
            preventScrollReset: true,
          });
        }}
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
        <GuideNavigation
          type="left"
          onClick={() =>
            setOffset((old) => previousPage(old, limit, filteredEpg.length))
          }
        />
        <GuideNavigation
          type="right"
          onClick={() =>
            setOffset((old) => nextPage(old, limit, filteredEpg.length))
          }
        />
      </>
    );
  };

  return (
    <div className={styles.container} ref={containerRef}>
      <div className={styles.header}>
        <div className={styles.bar}>
          <GuideControls
            search={search}
            onDayChange={handleDayChange}
            onSearch={(q) => {
              setSearch(q);
              setOffset(0);
              containerRef.current?.scrollTo(0, 0);
            }}
          />
        </div>
        <div className={c(styles.segment, styles.channels)}>
          {renderChannels()}
          {renderNavigation()}
        </div>
      </div>

      <div className={styles.segment}>{renderEventColumns()}</div>
    </div>
  );
}

export default GuideView;
