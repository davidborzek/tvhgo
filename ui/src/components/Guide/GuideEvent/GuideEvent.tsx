import moment from 'moment';
import { c } from '../../../utils/classNames';
import { parseTime } from '../../../utils/time';
import styles from './GuideEvent.module.scss';

type Props = {
  eventId: number;
  title: string;
  subtitle: string;
  description: string;
  startsAt: number;
  endsAt: number;
  showProgress?: boolean;
  onClick: (eventId: number) => void;
};

function renderTime(startsAt: number, endsAt: number): string {
  const minutesLeft = Math.floor((endsAt - startsAt) / 60);
  const start = parseTime(startsAt);
  const end = parseTime(endsAt);

  return `${start} - ${end} (${minutesLeft} Min.)`;
}

function GuideEvent({
  eventId,
  title,
  subtitle,
  description,
  startsAt,
  endsAt,
  showProgress,
  onClick,
}: Props) {
  const time = renderTime(startsAt, endsAt);
  const extra = subtitle || description;

  const renderProgress = () => {
    const width = Math.floor(
      ((moment().unix() - startsAt) / (endsAt - startsAt)) * 100
    );

    return (
      <div
        className={styles.progress}
        style={{
          width: `${width}%`,
        }}
      />
    );
  };

  return (
    <div className={styles.event} onClick={() => onClick(eventId)} tabIndex={0}>
      <span title={title} className={c(styles.name, styles.attribute)}>
        {title}
      </span>
      <span title={extra} className={c(styles.subtitle, styles.attribute)}>
        {extra}
      </span>
      <span title={time} className={c(styles.time, styles.attribute)}>
        {time}
      </span>
      {showProgress && renderProgress()}
    </div>
  );
}

export default GuideEvent;
