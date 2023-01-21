import { useTranslation } from 'react-i18next';
import { EpgEvent } from '../../../clients/api/api.types';
import { parseDatetime } from '../../../utils/time';
import Pair from '../../PairList/Pair/Pair';
import PairKey from '../../PairList/PairKey/PairKey';
import PairList from '../../PairList/PairList';
import PairValue from '../../PairList/PairValue/PairValue';
import EventRecordButton from '../EventRecordButton/EventRecordButton';
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
          <PairValue>{parseDatetime(event.startsAt, event.endsAt)}</PairValue>
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
