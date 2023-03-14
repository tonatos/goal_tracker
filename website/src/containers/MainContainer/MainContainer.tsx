import styles from './MainContainer.module.css';
import { Timeline } from '../../components/Timeline';
import { AddContribution } from '../../components/AddContribution';

export const MainContainer = () => {
    return (
        <div className={styles.App}>
            <div className={styles.Timeline}>
                <Timeline />
            </div>
            <div className={styles.AddContribution}>
                <AddContribution />
            </div>
        </div>
    )
}