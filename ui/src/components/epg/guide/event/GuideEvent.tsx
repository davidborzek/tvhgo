import Button from '@/components/common/button/Button';
import { c } from '@/utils/classNames';
import moment from 'moment';
import styles from './GuideEvent.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  eventId: number;
  title: string;
  subtitle: string;
  description: string;
  startsAt: number;
  endsAt: number;
  showProgress?: boolean;
  dvrState?: string;
  showDate?: boolean;
  channel?: string;
  onClick: (eventId: number) => void;
  onWatch?: () => void;
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
  showDate,
  channel,
  onClick,
  onWatch,
}: Props) {
  const { t } = useTranslation();

  const time = showDate
    ? t('event_datetime', { event: { startsAt, endsAt } })
    : t('event_time', { event: { startsAt, endsAt } });

  const extra = subtitle || description;

  const handleWatchClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onWatch?.();
  };

  const renderProgress = () => {
    const width = Math.floor(
      ((moment().unix() - startsAt) / (endsAt - startsAt)) * 100
    );

    if (width < 1) {
      return;
    }

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
      return (
        <span className={styles.recBadge} title={t('recording_running')} />
      );
    }
  };

  return (
    <div className={styles.event} onClick={() => onClick(eventId)} tabIndex={0}>
      <div className={styles.content}>
        <span title={title} className={c(styles.name, styles.attribute)}>
          {title}
        </span>
        <span title={extra} className={c(styles.subtitle, styles.attribute)}>
          {extra}
        </span>
        {channel && (
          <span title={channel} className={c(styles.channel, styles.attribute)}>
            {channel}
          </span>
        )}
        <span title={time} className={c(styles.time, styles.attribute)}>
          {time}
        </span>
      </div>
      {onWatch && (
        <div className={styles.actions}>
          <Button
            onClick={handleWatchClick}
            style="blue"
            size="small"
            label={t('watch')}
          />
        </div>
      )}
      {showProgress && renderProgress()}
      {renderRecBadge()}
    </div>
  );
}

export default GuideEvent;
