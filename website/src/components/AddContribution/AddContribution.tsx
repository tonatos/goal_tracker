import styles from './AddContribution.module.css';

export const AddContribution = ({}) => {
    return (
        <div className={styles.AddContribution}>
            <div className={styles.Input}>
                <input name="contribution" type="number" placeholder="0.00" />
            </div>
            <button className={styles.Button}>Внести</button>
        </div>
    )
}