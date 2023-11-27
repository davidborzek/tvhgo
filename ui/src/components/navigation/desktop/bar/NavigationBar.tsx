import { TvhgoHorizontalLogo } from '@/assets';
import NavigationItem from '../item/NavigationItem';

import styles from './NavigationBar.module.scss';
import { INavigationItem } from '../../types';

type Props = {
  items: INavigationItem[];
};

function NavigationBar({ items }: Props) {
  return (
    <div className={styles.root}>
      <div className={styles.head}>
        <TvhgoHorizontalLogo className={styles.logo} />
      </div>
      <div className={styles.items}>
        {items.map(({ icon, title, to }) => (
          <NavigationItem icon={icon} title={title} to={to} key={title} />
        ))}
      </div>
    </div>
  );
}

export default NavigationBar;
