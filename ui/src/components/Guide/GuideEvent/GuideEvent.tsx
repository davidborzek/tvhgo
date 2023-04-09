import moment from 'moment';
import { useTranslation } from 'react-i18next';
import { c } from '../../../utils/classNames';
import styles from './GuideEvent.module.scss';

type Props = {
  eventId: number;
  title: string;
  subtitle: string;
  description: string;
  startsAt: number;
  endsAt: number;
  showProgress?: boolean;
  dvrState?: string;
  onClick: (eventId: number) => void;
};

function GuideEvent({
  eventId,
  title,
  subtitle,
  description,
  startsAt,
  endsAt,
  showProgress,
  dvrState,
  onClick,
}: Props) {
  const { t } = useTranslation();

  const time = t('event_time', { event: { startsAt, endsAt } });
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

  const renderRecBadge = () => {
    if (dvrState === 'scheduled' || dvrState === 'recording') {
      return <span className={styles.recBadge} />;
    }
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
      {renderRecBadge()}
    </div>
  );
}

export default GuideEvent;
