import moment from 'moment';
import { c } from '../../../utils/classNames';
import styles from './GuideEvent.module.scss';

type Props = {
  title: string;
  subtitle: string;
  description: string;
  startsAt: number;
  endsAt: number;
  showProgress?: boolean;
};

function parseTime(ts: number): string {
  return new Date(ts * 1000).toLocaleTimeString(undefined, {
    hour: '2-digit',
    minute: '2-digit',
    hourCycle: 'h23',
  });
}

function renderTime(startsAt: number, endsAt: number): string {
  const minutesLeft = Math.floor((endsAt - startsAt) / 60);
  const start = parseTime(startsAt);
  const end = parseTime(endsAt);

  return `${start} - ${end} (${minutesLeft} Min.)`;
}

function GuideEvent({
  title,
  subtitle,
  description,
  startsAt,
  endsAt,
  showProgress,
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
    <div className={styles.event} tabIndex={0}>
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
