import moment from 'moment';
import { useTranslation } from 'react-i18next';
import Dropdown, { Option } from '../../Dropdown/Dropdown';
import Input from '../../Input/Input';
import styles from './GuideControls.module.scss';

type Props = {
  onSearch: (q: string) => void;
  onDayChange: (day: string) => void;
  search: string;
};

function GuideControls({ search, onSearch, onDayChange }: Props) {
  const { t } = useTranslation(['common']);

  const getDays = () => {
    const days: Option[] = [
      {
        title: t('today'),
        value: 'today',
      },
    ];

    for (let i = 1; i < 7; i++) {
      const date = moment().add(i, 'day').startOf('day');

      // TODO: date localization
      const title = `${t(`weekday_${date.day()}`)} (${date
        .toDate()
        .toLocaleDateString(undefined, {
          day: '2-digit',
          month: 'short',
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
      <Dropdown options={getDays()} onChange={onDayChange} />
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
