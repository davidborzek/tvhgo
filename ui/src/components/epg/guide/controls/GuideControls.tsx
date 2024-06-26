import Dropdown, { Option } from '@/components/common/dropdown/Dropdown';

import Input from '@/components/common/input/Input';
import moment from 'moment';
import styles from './GuideControls.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  onSearch: (q: string) => void;
  onDayChange: (day: string) => void;
  search: string;
  day: string;
};

function GuideControls({ day, search, onSearch, onDayChange }: Props) {
  const { t } = useTranslation();

  const getDays = () => {
    const days: Option[] = [
      {
        title: t('today'),
        value: 'today',
      },
    ];

    for (let i = 1; i < 7; i++) {
      const date = moment().add(i, 'day').startOf('day');

      const title = `${t(`weekday_${date.day()}`)} (${t('short_date', {
        ts: date.unix(),
      })})`;

      days.push({
        title,
        value: date.unix(),
      });
    }
    return days;
  };

  return (
    <div className={styles.controls}>
      <Dropdown value={day} options={getDays()} onChange={onDayChange} />
      <Input
        value={search}
        placeholder={t('search') || ''}
        onChange={(e) => {
          onSearch(e.target.value);
        }}
      />
    </div>
  );
}

export default GuideControls;
