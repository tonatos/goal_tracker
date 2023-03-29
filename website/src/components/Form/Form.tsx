import styles from './Form.module.css';
import CurrencyFormat from 'react-currency-format';
import Datetime, { DatetimepickerProps} from 'react-datetime';
import { Moment } from 'moment';
import 'moment/locale/ru';
import "react-datetime/css/react-datetime.css";
import { IoCloseCircle } from "react-icons/io5";
import { ChangeEvent, FormEvent, ReactElement } from 'react';


export enum InputType {
    Text = 'text',
    Currency = 'currency',
    Date = 'date',
}
export enum ButtonType {
    Ok = 'ok',
    Cancel = 'cancel',
}

interface InputProp {
    name: string,
    value?: string | number | Date,
    label?: string,
    large?: boolean,
    type?: InputType,
    placeholder?: string,
    icon?: JSX.Element,
    inputProps?: object,
    onChange?: (props: {name: string, value: string | number | Date}) => void,
}

const BaseInput = (props: {
    large: boolean,
    passProps: object,
    value?: string,
    onChange: () => void
}) => (
    <input
        type={'text'}
        value={props.value ? props.value : ''}
        onChange={props.onChange}
        {...props.passProps}
        className={`${props?.large ? styles.InputInputLarge : styles.InputInput}`}
    />
);

const TextInput = (props: any) => (
    <BaseInput
        {...{
            passProps: {
                ...props.passProps,
                value: props.value
            },
            large: props.large,
            onChange: props.onChange,
        }}
    />
);

const CurrencyInput = (props: any) => (
    <CurrencyFormat
        value={props.value ? props.value : ''}
        customInput={BaseInput}
        thousandSeparator={' '}
        onValueChange={props.onChange}
    />
);

const DateInput = (props: any) => (
    <Datetime
        locale="ru"
        timeFormat={false}
        initialValue={props.value ? new Date(props.value) : ''}
        renderInput={(datetimeProps: DatetimepickerProps, openCalendar, closeCalendar) => (
            BaseInput({
                large: props.large,
                passProps: {
                    ...props.passProps,
                    ...datetimeProps,
                },
                onChange: () => {}
            })
        )}
        onChange={props.onChange}
    />
);

export const Input = (props: InputProp) => {
    const id: string = `id-${props.name}`;
    const RenderInput = () => {
        const inputProps = {
            large: props.large,
            value: props.value,
            passProps: {
                id: id,
                name: props.name,
                placeholder: props.placeholder,
                ...(props.inputProps ? props.inputProps : {}),
            }
        }
        switch (props.type) {
            case 'date':
                return DateInput({
                    ...inputProps,
                    onChange: (props.onChange ? (e: Moment) => {
                        if (props.onChange) {
                            props.onChange({
                                name: props.name,
                                value: e && e.toDate ? e.toDate() : new Date(),
                            })
                        }
                    } : () => {}),
                });

            case 'currency':
                return CurrencyInput({
                    ...inputProps,
                    onChange: (props.onChange ? (value: CurrencyFormat.Values) => {
                        if (props.onChange) {
                            props.onChange({
                                name: props.name,
                                value: value.floatValue
                            })
                        }
                    } : () => {}),
                });

            default:
            case 'text':
                return TextInput({
                    ...inputProps,
                    onChange: (props.onChange ? (e: FormEvent<HTMLInputElement>) => {
                        if (props.onChange) {
                            props.onChange({
                                name: e.currentTarget.name,
                                value: e.currentTarget.value
                            })
                        }
                    } : () => {}),
                });
        }
    }
    
    return (
        <div className={
            [
                styles.Input,
                props.type ? styles[`Input--${props.type}`] : styles[`Input-text`],
                props.icon ? styles[`Input--Icon`] : '',
            ].join(' ')
        }>
            {props.label ? <label className={styles.InputLabel} htmlFor={id}>{props.label}</label>: <></>}
            {props.icon ? <div className={styles.InputIcon}>{props.icon}</div>: <></>}
            {RenderInput()}
        </div>
    )
}


interface ButtonProp {
    text: string,
    type?: ButtonType,
    icon?: ReactElement,
    buttonProps?: object,
}

export const Button = (props: ButtonProp) => {
    return (
        <div className={props.type == ButtonType.Cancel ? styles.ButtonCancel : styles.Button}>
            <button className={styles.ButtonButton} {...props.buttonProps}>
                {props.icon ? <span className={styles.ButtonIcon}>{props.icon}</span> : <></>}
                {props.text}
            </button>
        </div>
    );
};