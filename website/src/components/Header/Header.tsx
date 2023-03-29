import { useState } from 'react';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import { useGoalActions } from '../../actions/goal.actions';
import { ModalState } from '../../interfaces/modal.interfaces';

import { currentGoalAtom, goalsListQuery } from '../../recoil/goal';
import { modalStateAtom } from '../../recoil/modal';
import { Popover } from '../Popover';
import styles from './Header.module.css';

interface HeaderProp {
    createGoalModal: any
}

export const Header = (prop: HeaderProp ) => {
    const goalActions = useGoalActions();
    const [ popoverShow, setPopoverShow ] = useState<boolean>(false);
    const goals = useRecoilValue(goalsListQuery);
    const [ currentGoal, setCurrentGoal ] = useRecoilState(currentGoalAtom);
    const [ modalState, setModalState ] = useRecoilState(modalStateAtom);
   
    const popoverOnItemClick = (e: React.MouseEvent, id: number) => {
        const selectedGoalId = goals.filter(el => el.id == id)[0].id;
        goalActions.getDefault(selectedGoalId);
    }

    const popoverOnMenuItemClick = (e: React.MouseEvent, id: number) => {
        e.stopPropagation();
        setModalState({
            createUpdateGoal: {
                state: true,
                params: {
                    goal: goals.filter(el => el.id == id)[0]
                }
            }
        } as ModalState)
    }

    return (
        <div className={styles.Header}>
            {Object.keys(currentGoal).length ? (
                <div className={styles.HeaderSelectorContainer}>
                    Сейчас стремимся к:&nbsp;
                    <a
                        href="#"
                        className={styles.Selector}
                        onClick={() => setPopoverShow(!popoverShow)}
                    >
                        {currentGoal.name}
                    </a>
                    
                    <div className={styles.HeaderSelectorPopoverContainer}>
                        <Popover
                            showState={[popoverShow, setPopoverShow]}
                            items={goals.map(item => ({id: item.id ? item.id : 0, value: item.name}))}
                            currentItem={currentGoal.id ? currentGoal.id : 0}
                            onItemClick={popoverOnItemClick}
                            onMenuItemClick={popoverOnMenuItemClick}
                        />
                    </div>
                </div>
            ) : (
                <>Пока целей не добавлено.&nbsp;
                <a
                    href="#"
                    className={styles.CreateButton}
                    onClick={prop.createGoalModal}
                >Создать цель</a>?</>
            )}
        </div>
    )
}