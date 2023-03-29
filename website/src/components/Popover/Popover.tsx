import { Dispatch, SetStateAction } from 'react';
import { AiOutlinePlusCircle } from 'react-icons/ai';
import { FiMoreHorizontal } from 'react-icons/fi';
import { useRecoilState } from 'recoil';
import { ModalState } from '../../interfaces/modal.interfaces';
import { modalStateAtom } from '../../recoil/modal';
import { Button, ButtonType } from '../Form';
import styles from './Popover.module.css';

interface PopoverProps {
    showState?: [boolean, Dispatch<SetStateAction<boolean>>],
    items: Array<{id: number, value: string}>,
    currentItem: number,
    onItemClick: (e: React.MouseEvent, id: number) => void,
    onMenuItemClick: (e: React.MouseEvent, id: number) => void,
}

export const Popover = (props: PopoverProps) => {
    const [ modalState, setModalState ] = useRecoilState(modalStateAtom);
    const [ show, setShow ] = props.showState ? props.showState : [false, () => {}];

    const Popover = (
        <div className={styles.Popover}>
            <ul className={styles.PopoverList}>
                {props.items.map(item => (
                    <li
                        className={[
                            styles.PopoverListItem,
                            item.id == props.currentItem ? styles.PopoverListItemActive : ''
                        ].join(' ')}
                        key={`item-${item.id}`}
                        data-id={item.id}
                        onClick={e => props.onItemClick(e, item.id)}
                    >
                        <span className={styles.PopoverListItemValue}>{item.value}</span>
                        <span
                            className={styles.PopoverListItemMenu}
                            onClick={e => props.onMenuItemClick(e, item.id)}
                        >
                            <FiMoreHorizontal size={24} />
                        </span>
                    </li>
                ))}
            </ul>
            <Button
                text={'Добавить цель'}
                type={ButtonType.Ok}
                icon={<AiOutlinePlusCircle size={24} />}
                buttonProps={{
                    onClick: () => {
                        setModalState({
                            ...modalState,
                            createUpdateGoal: {state: true}
                        } as ModalState);
                    }
                }}
            />
        </div>
    );
    return show ? Popover : <></>;
}