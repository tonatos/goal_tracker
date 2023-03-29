import { useEffect } from 'react';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Header } from '../../components/Header';
import { Modal } from '../../components/Modal';
import { Timeline } from '../../components/Timeline';
import { AddContribution } from '../../components/AddContribution';

import styles from './MainContainer.module.css';

import { useGoalActions } from '../../actions/goal.actions';
import { goalsListQuery, currentGoalAtom } from '../../recoil/goal';
import { modalStateAtom } from '../../recoil/modal';
import { ModalState } from '../../interfaces/modal.interfaces';
import { GoalForm } from '../../components/GoalForm';
import { ToastContainer } from 'react-toastify';

const CreateGoalModal = () => {
    const [ modalState, setModalState ] = useRecoilState<ModalState>(modalStateAtom);
    return (
        <Modal
            isShow={modalState?.createUpdateGoal?.state || false}
            closeModal={() => setModalState({ createUpdateGoal: {state: false} } as ModalState)}
        >
            <GoalForm
                {...(modalState.createUpdateGoal?.params?.goal ? {
                    goal: modalState.createUpdateGoal.params.goal
                } : {})}
            />
        </Modal>
    )
}

const ConstributionListModal = () => {
    const [ modalState, setModalState ] = useRecoilState<ModalState>(modalStateAtom);
    return (
        <Modal
            isShow={modalState?.listConstribution?.state || false}
            closeModal={() => setModalState({ listConstribution: {state: false} } as ModalState)}
            header={'Список взносов'}>
            Modal Contribtion!
        </Modal>
    )
}

export const MainContainer = () => {
    const [ modalState, setModalState ] = useRecoilState(modalStateAtom);
    const goals = useRecoilValue(goalsListQuery);
    const goalActions = useGoalActions();
    const currentGoal = useRecoilValue(currentGoalAtom);

    useEffect(() => {
        goalActions.getAll();
    }, []);

    useEffect(() => {
        const currentId = [...goals].sort((a,b) => a.id && b.id ? a.id - b.id : -1)[0]?.id;
        goalActions.getDefault(currentId);
    }, [goals]);

    return (
        <div>
            <ToastContainer theme="colored" />

            <CreateGoalModal />
            <ConstributionListModal />

            <div className={styles.App}>
                <div className={styles.Header}>
                    <Header createGoalModal={() => setModalState({createUpdateGoal: {state: true}} as ModalState)} />
                </div>
                {Object.keys(currentGoal).length ? (
                    <>
                        <div className={styles.Timeline}>
                            <Timeline currentGoal={currentGoal} />
                        </div>
                        <div className={styles.AddContribution}>
                            <AddContribution />
                        </div>
                    </>
                ) : (
                    <div className={styles.EmptyGoal}>
                        <div>
                            <a
                                href="#"
                                className={styles.EmptyGoalAdd}
                                onClick={() => setModalState({createUpdateGoal: {state: true}} as ModalState)}
                            >Добавьте цель</a>,
                            чтобы следить за ее прогрессом!
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}