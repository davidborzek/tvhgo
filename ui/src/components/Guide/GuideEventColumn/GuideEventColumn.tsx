import { EpgEvent } from "../../../clients/api/api.types";
import GuideEvent from "../GuideEvent/GuideEvent";
import styles from "./GuideEventColumn.module.scss";

type Props = {
  events: EpgEvent[];
};

function GuideEventColumn({ events }: Props) {
  const renderEvents = () => {
    return events.map((event, index) => (
      <GuideEvent
        key={event.id}
        title={event.title}
        description={event.description}
        subtitle={event.subtitle}
        startsAt={event.startsAt}
        endsAt={event.endsAt}
        showProgress={!index}
      />
    ));
  };

  return <div className={styles.column}>{renderEvents()}</div>;
}

export default GuideEventColumn;
