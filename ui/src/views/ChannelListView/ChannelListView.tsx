import React, { useCallback } from 'react';
import { useFetchEpg } from '../../hooks/epg';

import styles from './ChannelListView.module.scss';
import { useTranslation } from 'react-i18next';
import ChannelListItem from '../../components/ChannelListItem/ChannelListItem';
import { GetEpgEventsQuery } from '../../clients/api/api';
import Error from '../../components/Error/Error';
import { useLoading } from '../../contexts/LoadingContext';

const limit = 50;

const opts: GetEpgEventsQuery = {
  nowPlaying: true,
  limit,
  sort_key: 'channelNumber',
  sort_dir: 'asc',
};

function ChannelListView() {
  const { t } = useTranslation();
  const { events, offset, total, error, increaseOffset } = useFetchEpg(opts);

  const { isLoading } = useLoading();

  const handleScroll = useCallback<React.UIEventHandler<HTMLDivElement>>(
    (evt) => {
      const { scrollHeight, scrollTop, clientHeight } = evt.currentTarget;

      if (
        scrollHeight - scrollTop - clientHeight < 50 &&
        total > offset &&
        !isLoading
      ) {
        increaseOffset(limit);
      }
    },
    [total, isLoading, offset]
  );

  const renderChannels = () => {
    return events.map((event) => {
      return <ChannelListItem key={event.id} event={event} />;
    });
  };

  if (error) {
    return <Error message={error} />;
  }

  return (
    <div className={styles.container} onScroll={handleScroll}>
      <div className={styles.channelList}>
        <h1>{t('channels')}</h1>
        {renderChannels()}
      </div>
    </div>
  );
}

export default ChannelListView;
