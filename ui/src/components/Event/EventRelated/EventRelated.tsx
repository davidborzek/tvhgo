import { Link } from 'react-router-dom';
import { EpgEvent } from '../../../clients/api/api.types';
import styles from './EventRelated.module.scss';

type Props = {
  relatedEvents: EpgEvent[];
};

function renderDatetime(ts: number) {
  return new Date(ts * 1000).toLocaleString(undefined, {
    hour: '2-digit',
    minute: '2-digit',
    month: '2-digit',
    day: '2-digit',
  });
}

function EventRelated({ relatedEvents }: Props) {
  const renderRelatedEvents = () => {
    return relatedEvents.map((event) => {
      return (
        <Link
          className={styles.relatedEvent}
          key={event.id}
          to={`/guide/events/${event.id}`}
        >
          {renderDatetime(event.startsAt)} - {event.title}
        </Link>
      );
    });
  };

  if (relatedEvents.length == 0) {
    return <></>;
  }

  return (
    <div className={styles.EventRelated}>
      <h2>Related Events</h2>
      {renderRelatedEvents()}
    </div>
  );
}

export default EventRelated;
