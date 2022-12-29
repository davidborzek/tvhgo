import { useFetchChannelEvents } from "../../hooks/epg";
import styles from "./Guide.module.scss";
import moment from "moment";
import { useEffect, useRef, useState } from "react";
import GuideChannel from "../../components/Guide/GuideChannel/GuideChannel";
import GuideEventColumn from "../../components/Guide/GuideEventColumn/GuideEventColumn";
import GuideNavigation from "../../components/Guide/GuideNavigation/GuideNavigation";
import { EpgChannel } from "../../clients/api/api.types";
import GuideControls from "../../components/Guide/GuideControls/GuideControls";

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

function prepareEpg(
  epg: EpgChannel[],
  search: string,
  offset: number,
  limit: number
) {
  return epg
    .filter((e) =>
      e.channelName.toLowerCase().includes(search.toLocaleLowerCase())
    )
    .slice(offset, limit + offset);
}

function Guide() {
  const containerRef = useRef<HTMLDivElement>(null);

  const { events, setStartsAt, setEndsAt } = useFetchChannelEvents({
    endsAt: moment().add(24, "hour").unix(),
    sort_key: "channelNumber",
    limit: 100,
  });

  const setDate = (start?: number, end?: number) => {
    setStartsAt(start);
    setEndsAt(end);
    containerRef.current?.scrollTo(0, 0);
  };

  const [search, setSearch] = useState("");

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

      if (innerWidth > 1200 && current !== 5) {
        setLimit(5);
        return;
      }

      if (innerWidth > 800 && innerWidth <= 1200 && current !== 4) {
        setLimit(4);
        return;
      }
      if (innerWidth > 500 && innerWidth <= 800 && current !== 3) {
        setLimit(3);
        return;
      }

      if (innerWidth <= 500 && current !== 2) {
        setLimit(2);
        return;
      }
    };

    limitResults();
    window.addEventListener("resize", limitResults);

    return () => {
      window.removeEventListener("resize", limitResults);
    };
  }, []);

  const preparedEpg = prepareEpg(events, search, offset, limit);

  const handleDayChange = (dateString: string) => {
    if (dateString === "today") {
      setDate(undefined, moment().endOf("day").unix());
      return;
    }

    const date = moment.unix(parseInt(dateString, 10));
    setDate(date.unix(), date.endOf("day").unix());
  };

  const renderChannels = () => {
    return preparedEpg.map((event) => (
      <GuideChannel
        key={event.channelId}
        name={event.channelName}
        picon={`/api/picon/${event.piconId}`}
        number={event.channelNumber}
      />
    ));
  };

  const renderEventColumns = () => {
    return preparedEpg.map((channel) => (
      <GuideEventColumn key={channel.channelId} events={channel.events} />
    ));
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
              containerRef.current?.scrollTo(0, 0);
            }}
          />
        </div>
        <div className={styles.segment}>{renderChannels()}</div>
      </div>

      <div className={styles.segment}>{renderEventColumns()}</div>

      <GuideNavigation
        type="left"
        onClick={() =>
          setOffset((old) => previousPage(old, limit, events.length))
        }
      />
      <GuideNavigation
        type="right"
        onClick={() => setOffset((old) => nextPage(old, limit, events.length))}
      />
    </div>
  );
}

export default Guide;
