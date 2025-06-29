import { EpgEvent } from '@/clients/api/api.types';
import Button from '@/components/common/button/Button';
import EventRecordButton from '../recordButton/EventRecordButton';
import Pair from '@/components/common/pairList/Pair/Pair';
import PairKey from '@/components/common/pairList/PairKey/PairKey';
import PairList from '@/components/common/pairList/PairList';
import PairValue from '@/components/common/pairList/PairValue/PairValue';
import styles from './EventInfo.module.scss';
import { useTranslation } from 'react-i18next';
import moment from 'moment';

type Props = {
  event: EpgEvent;
  pending: boolean;
  handleOnRecord: () => void;
  onWatch?: () => void;
};

function EventInfo({ event, handleOnRecord, pending, onWatch }: Props) {
  const { t } = useTranslation();

  const isCurrentlyPlaying = () => {
    const now = moment().unix();
    return now >= event.startsAt && now < event.endsAt;
  };

  return (
    <div className={styles.EventInfo}>
      <h1>{event.title}</h1>
      <div>
        {isCurrentlyPlaying() && onWatch && (
          <Button
            onClick={onWatch}
            style="blue"
            size="small"
            label={t('watch')}
          />
        )}
        <EventRecordButton
          pending={pending}
          onClick={handleOnRecord}
          dvrUuid={event.dvrUuid}
        />
      </div>
      <PairList>
        <Pair>
          <PairKey>{t('subtitle')}</PairKey>
          <PairValue>{event.subtitle}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('airs')}</PairKey>
          <PairValue>{t('event_datetime', { event })}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('description')}</PairKey>
          <PairValue>{event.description}</PairValue>
        </Pair>
      </PairList>
    </div>
  );
}

export default EventInfo;
