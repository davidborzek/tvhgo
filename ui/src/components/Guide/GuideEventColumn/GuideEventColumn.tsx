import { useNavigate, useNavigation } from 'react-router-dom';
import { EpgEvent } from '../../../clients/api/api.types';
import GuideEvent from '../GuideEvent/GuideEvent';
import styles from './GuideEventColumn.module.scss';

type Props = {
  events: EpgEvent[];
};

function GuideEventColumn({ events }: Props) {
  const navigate = useNavigate();

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
        onClick={(id) => {
          navigate(`/guide/events/${id}`);
        }}
      />
    ));
  };

  return <div className={styles.column}>{renderEvents()}</div>;
}

export default GuideEventColumn;
