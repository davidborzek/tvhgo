import { EpgEvent } from '@/clients/api/api.types';
import styles from './GuideEventColumn.module.scss';
import GuideEvent from '../event/GuideEvent';

type Props = {
  events: EpgEvent[];
  onClick: (eventId: number) => void;
};

function GuideEventColumn({ events, onClick }: Props) {
  const renderEvents = () => {
    return events.map((event, index) => (
      <GuideEvent
        eventId={event.id}
        key={event.id}
        title={event.title}
        description={event.description}
        subtitle={event.subtitle}
        startsAt={event.startsAt}
        endsAt={event.endsAt}
        showProgress={!index}
        dvrState={event.dvrState}
        onClick={onClick}
      />
    ));
  };

  return <div className={styles.column}>{renderEvents()}</div>;
}

export default GuideEventColumn;
