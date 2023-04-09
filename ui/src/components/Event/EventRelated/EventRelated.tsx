import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';
import { EpgEvent } from '../../../clients/api/api.types';
import Pair from '../../PairList/Pair/Pair';
import PairKey from '../../PairList/PairKey/PairKey';
import PairList from '../../PairList/PairList';
import PairValue from '../../PairList/PairValue/PairValue';
import styles from './EventRelated.module.scss';

type Props = {
  relatedEvents: EpgEvent[];
};

function EventRelated({ relatedEvents }: Props) {
  const { t } = useTranslation();

  const renderTitle = (event: EpgEvent) => {
    const datetime = t('event_datetime', { event });
    const subtitle = event.subtitle ? ` • ${event.subtitle}` : '';
    return `${datetime}${subtitle}`;
  };

  const renderRelatedEvents = () => {
    return relatedEvents.map((event) => {
      return (
        <PairList>
          <Pair>
            <PairKey>{event.channelName}</PairKey>
            <PairValue>
              <Link
                className={styles.link}
                key={event.id}
                to={`/guide/events/${event.id}`}
              >
                {renderTitle(event)}
              </Link>
            </PairValue>
          </Pair>
        </PairList>
      );
    });
  };

  if (relatedEvents.length == 0) {
    return <></>;
  }

  return (
    <div className={styles.EventRelated}>
      <h2>{t('related_events')}</h2>
      {renderRelatedEvents()}
    </div>
  );
}

export default EventRelated;
