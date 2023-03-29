import { Input, InputType, Button, ButtonType } from '../Form';
import { ImageUpload } from '../ImageUpload';
import styles from './GoalForm.module.css';
import { IoCalendarOutline, IoCloseCircle } from "react-icons/io5";
import { FormEvent, useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import { CreateUpdateGoal, Goal } from '../../interfaces/goal.interfaces';
import { useImageUploadActions } from '../../actions/upload.actions';
import { useGoalActions } from '../../actions/goal.actions';
import { useRecoilState, useRecoilValue } from 'recoil';
import { modalStateAtom } from '../../recoil/modal';
import { ModalState } from '../../interfaces/modal.interfaces';
import { goalsListQuery } from '../../recoil/goal';

interface GoalFormProps {
    goal?: Goal
}

export const GoalForm = (props: GoalFormProps) => {
    const goalActions = useGoalActions();
    const uploadImageAction = useImageUploadActions();
    const goals = useRecoilValue(goalsListQuery);
    
    const [ modalState, setModalState ] = useRecoilState(modalStateAtom);
    const [ submit, setSubmit ] = useState<boolean>(false);
    const [ createSuccess, setCreateSuccess ] = useState<boolean>(false);
    const [ deleteSuccess, setDeleteSuccess ] = useState<boolean>(false);
    const [ error, setError ] = useState<string>();
    const [ uploadedImage, setUploadedImage ] = useState<File>();
    const [ goalForm, setGoalForm ] = useState<CreateUpdateGoal>(props.goal as CreateUpdateGoal || {
        name: '',
        goal_amount: 0,
        target_date: undefined,
        image: '',
    });

    const inputOnChange = (props: {name: string, value: string | number | Date}) => {
        setGoalForm({
            ...goalForm,
            [props.name]: props.value,
        });
    }

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        setSubmit(true);

        if (uploadedImage) {
            const images = await uploadImageAction.upload(uploadedImage);
            setGoalForm({
                ...goalForm,
                image: images['image']
            });
        }
    }

    useEffect(() => {
        if (
            !goalForm.image ||
            !submit
        ) {
            return;
        }
        try {
            goalActions.create(goalForm);
            setCreateSuccess(true);

            toast.success('Цель успешно создана!', {autoClose: 2000});
            toast.onChange((notify) => {
                if (notify.status == 'removed') {
                    setModalState({createUpdateGoal: {state: false}} as ModalState);
                }
            });
        } catch (e) {
            setError(e as string);
        }
    }, [goalForm]);

    useEffect(() => {
        if (!deleteSuccess) return;
        goalActions.getDefault(goals[0].id);
    }, [goals])

    const handleDelete = async (e: MouseEvent) => {
        e.preventDefault();
        try {
            await goalActions.del(goalForm.id);
            setDeleteSuccess(true);
            toast.success('Убрали лишнюю цель', {autoClose: 2000});
            setModalState({createUpdateGoal: {state: false}} as ModalState);
        } catch (e: any) {
            toast.error(e.message, {autoClose: 2000});
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <div className={styles.FormRow}>
                <Input
                    large={true}
                    label={'К чему стремимся?'}
                    placeholder={'Назовите вашу цель'}
                    onChange={inputOnChange}
                    name={'name'}
                    value={goalForm.name}
                    inputProps={{required: true}}
                />
            </div>
            <div className={styles.FormRow}>
                <Input
                    label={'Сколько надо денег:'}
                    type={InputType.Currency}
                    placeholder={'0.00'}
                    onChange={inputOnChange}
                    name={'goal_amount'}
                    value={goalForm.goal_amount}
                    inputProps={{required: true}}
                />
                <Input
                    label={'К какому сроку:'}
                    type={InputType.Date}
                    icon={<IoCalendarOutline size={'20px'}/>}
                    onChange={inputOnChange}
                    name={'target_date'}
                    value={goalForm.target_date}
                    inputProps={{required: true, autoComplete: 'off'}}
                />
            </div>
            <div className={styles.FormRow}>
                <ImageUpload
                    backgroundText={goalForm.name}
                    inputProps={{required: true}}
                    onUpload={(file: File) => setUploadedImage(file)}
                    initial={goalForm.image ? goalForm.image : undefined}
                />
            </div>
            
            <div className={styles.FormRow}>
                <Button text={'Сохранить'} buttonProps={{
                    type: 'submit',
                    disabled: createSuccess,
                }} />
                {props.goal ? <Button
                    text={'Удалить цель'}
                    type={ButtonType.Cancel}
                    icon={<IoCloseCircle size={24} />}
                    buttonProps={{
                        onClick: handleDelete,
                        disabled: deleteSuccess
                    }}
                /> : <></>}
            </div>
            
            {error ? <div className={styles.FormRow}>
                <div className={styles.Error}>
                    <h3>При создании цели случились ошибки</h3>
                    {error}
                </div>
            </div> : <></>}
        </form>
    )
}