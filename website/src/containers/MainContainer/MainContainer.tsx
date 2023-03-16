import styles from './MainContainer.module.css';
import { Timeline } from '../../components/Timeline';
import { AddContribution } from '../../components/AddContribution';
import { useEffect } from 'react';
import { useGoalActions } from '../../actions/goal.actions';
import { useRecoilValue } from 'recoil';
import { goalsListQuery } from '../../recoil/goal';

export const MainContainer = () => {
    const goalActions = useGoalActions();
    const goals = useRecoilValue(goalsListQuery);

    useEffect(() => {
        goalActions.getAll();
    }, []);

    useEffect(() => {
        const currentId = [...goals].sort((a,b) => a.id && b.id ? a.id - b.id : -1)[0]?.id;
        goalActions.getDefault(currentId);
    }, [goals]);

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