import { useEffect, useState } from 'react';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import { currentGoalAtom } from '../../recoil/goal';
import Contribution from "../../interfaces/contribution.interfaces";
import styles from './AddContribution.module.css';
import { useContributionActions } from '../../actions/contribution.actions';
import { useGoalActions } from '../../actions/goal.actions';
import { contributionsListQuery } from '../../recoil/contribution';

export const AddContribution = ({}) => {
    const [ showError, setShowError ] = useState(false)
    const [ addContributionAmount, setAddContributionAmount ] = useState(0);
    const [ currentGoal, setCurrentGoal] = useRecoilState(currentGoalAtom);
    const constributions = useRecoilValue(contributionsListQuery);
    const contributionActions = useContributionActions();
    const goalsActions = useGoalActions();

    const addContribution = async () => {
        if (!addContributionAmount) {
            setShowError(true);
            setTimeout(() => setShowError(false), 2000)
            return;
        }
        setShowError(false);
        
        await contributionActions.create(currentGoal.id, {amount: addContributionAmount})
        return false;
    }
    useEffect(() => {
        if (constributions.length == 0) {
            return;
        }
        goalsActions.getDefault(currentGoal.id)
    }, [constributions])

    return (
        <div className={`${styles.AddContribution} ${showError ? styles.Error : ''}`}>
            <form className={styles.Form} onSubmit={(e) => {addContribution(); e.preventDefault();}}>
                <div className={styles.Input}>
                    
                    <input
                        name="contribution"
                        type="number"
                        placeholder="0.00"
                        onChange={(e) => setAddContributionAmount(e.currentTarget.valueAsNumber)}
                    />
                </div>
                <button className={styles.Button} type="submit">Внести</button>
            </form>
        </div>
    )
}