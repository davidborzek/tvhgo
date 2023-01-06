import { useTranslation } from 'react-i18next';
import { EpgEvent } from '../../../clients/api/api.types';
import { parseTime } from '../../../utils/time';
import EventRecordButton from '../EventRecordButton/EventRecordButton';
import styles from './EventInfo.module.scss';

type Props = {
  event: EpgEvent;
  pending: boolean;
  handleOnRecord: () => void;
};

function EventInfo({ event, handleOnRecord, pending }: Props) {
  const { t } = useTranslation();

  function renderDatetime(startsAt: number, endsAt: number): string {
    const date = new Date(startsAt * 1000);
    const day = date.getDay();

    const minutesLeft = Math.floor((endsAt - startsAt) / 60);
    const start = parseTime(startsAt);

    return `${t(
      `weekday_${day}`
    )}, ${date.toLocaleDateString()} • ${start} • ${minutesLeft} Min.`;
  }

  return (
    <div className={styles.EventInfo}>
      <h1>{event.title}</h1>
      <div>
        <EventRecordButton
          pending={pending}
          onClick={handleOnRecord}
          dvrUuid={event.dvrUuid}
        />
      </div>
      <span>{renderDatetime(event.startsAt, event.endsAt)}</span>
      <span>{event.subtitle}</span>
      <span>{event.description}</span>
    </div>
  );
}

export default EventInfo;
