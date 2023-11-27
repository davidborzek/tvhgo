import { useTranslation } from 'react-i18next';
import { EpgEvent } from '@/clients/api/api.types';
import Pair from '@/components/common/pairList/Pair/Pair';
import PairKey from '@/components/common/pairList/PairKey/PairKey';
import PairList from '@/components/common/pairList/PairList';
import PairValue from '@/components/common/pairList/PairValue/PairValue';
import EventRecordButton from '../recordButton/EventRecordButton';
import styles from './EventInfo.module.scss';

type Props = {
  event: EpgEvent;
  pending: boolean;
  handleOnRecord: () => void;
};

function EventInfo({ event, handleOnRecord, pending }: Props) {
  const { t } = useTranslation();

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
