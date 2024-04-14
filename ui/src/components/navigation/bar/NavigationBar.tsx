import { INavigationItem } from '../types';
import NavigationItem from '../item/NavigationItem';
import styles from './NavigationBar.module.scss';

type Props = {
  items: INavigationItem[];
};

function NavigationBar({ items }: Props) {
  return (
    <div className={styles.root}>
      <div className={styles.items}>
        {items.map(({ icon, title, to, items: children }) => (
          <NavigationItem
            topLevel
            icon={icon}
            title={title}
            to={to}
            key={title}
            items={children}
          />
        ))}
      </div>
    </div>
  );
}

export default NavigationBar;
