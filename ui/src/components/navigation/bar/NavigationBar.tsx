import { INavigationItem } from '../types';
import NavigationItem from '../item/NavigationItem';
import styles from './NavigationBar.module.scss';

type Props = {
  items: INavigationItem[];
  roles?: string[];
};

function NavigationBar({ items, roles = [] }: Props) {
  return (
    <div className={styles.root}>
      <div className={styles.items}>
        {items.map(({ icon, title, to, items: children, requiredRoles }) => (
          <NavigationItem
            topLevel
            icon={icon}
            title={title}
            to={to}
            key={title}
            items={children}
            requiredRoles={requiredRoles}
            roles={roles}
          />
        ))}
      </div>
    </div>
  );
}

export default NavigationBar;
