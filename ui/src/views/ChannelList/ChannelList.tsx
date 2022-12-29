import React, { useCallback } from 'react';
import { useFetchEpg } from '../../hooks/epg';

import styles from './ChannelList.module.scss';
import { useTranslation } from 'react-i18next';
import ChannelListItem from '../../components/ChannelListItem/ChannelListItem';
import { GetEpgEventsQuery } from '../../clients/api/api';

const limit = 50;

const opts: GetEpgEventsQuery = {
  nowPlaying: true,
  limit,
  sort_key: 'channelNumber',
  sort_direction: 'asc',
};

function ChannelList() {
  const { t } = useTranslation('menu');
  const { events, offset, total, loading, increaseOffset } = useFetchEpg(opts);

  const handleScroll = useCallback<React.UIEventHandler<HTMLDivElement>>(
    (evt) => {
      const { scrollHeight, scrollTop, clientHeight } = evt.currentTarget;

      if (
        scrollHeight - scrollTop - clientHeight < 50 &&
        total > offset &&
        !loading
      ) {
        increaseOffset(limit);
      }
    },
    [total, loading, offset]
  );

  const renderChannels = () => {
    return events.map((event) => {
      return <ChannelListItem event={event} />;
    });
  };

  return (
    <div className={styles.container} onScroll={handleScroll}>
      <div className={styles.channelList}>
        <h1>{t('channels')}</h1>
        {renderChannels()}
      </div>
    </div>
  );
}

export default ChannelList;
