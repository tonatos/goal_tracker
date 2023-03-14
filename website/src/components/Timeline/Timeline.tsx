import StartPointImage from '../../assets/StartPoint.png';
import FinishPointImage from '../../assets/FinishPoint.png';
import CursorIcon from '../../assets/CursorIcon.svg';
import styles from './Timeline.module.css'

import { useState, useRef, useEffect, useLayoutEffect } from 'react';
import CurrencyFormat from 'react-currency-format';



interface CursorProp {
    percent: number,
    containerWidth: number,
    leftOffset: number
}

interface FinishPoint {
    goalAmount: number,
    containerWidth: number,
    leftOffset: number
}

const StartPoint = ({}) => {
    return (
        <div className={styles.StartPoint}>
            <img src={StartPointImage} className={styles.Image} />
        </div>
    )
}

const FinishPoint = (prop: FinishPoint) => {
    const goalAmountRef = useRef<SVGTextElement | null>(null);
    const calculateGradientPoint = (): number => {
        const goalAmountRect = goalAmountRef?.current?.getBoundingClientRect() || {} as DOMRect;
        const cursorPosition = prop.containerWidth - prop.leftOffset;
        
        if (cursorPosition < goalAmountRect.width) {
            return (goalAmountRect.width - cursorPosition) / (goalAmountRect.width / 100) / 100;
        } else {
            return 0;
        }
    }
    return (
        <div className={styles.FinishPoint}>
            <img src={FinishPointImage} className={styles.Image} />
            <div className={styles.FinishAmountContainer}>
                <svg xmlns="http://www.w3.org/2000/svg" width="340" height="90">
                    <defs>
                        <linearGradient id="gradient" x1="-0.04" x2="1.04">
                            <stop stopColor="#0DCFAA" offset="0"/>
                            <stop stopColor="#0DCFAA" offset={calculateGradientPoint()}/>
                            <stop stopColor="#ffffff" offset="0"/>
                            <stop stopColor="#ffffff" offset="1"/>
                        </linearGradient>
                    </defs>
                    <text
                        ref={goalAmountRef}
                        textAnchor="middle"
                        x="50%" y="50%"
                        className={styles.FinishAmount}>
                        <CurrencyFormat
                            value={prop.goalAmount}
                            displayType={'text'}
                            thousandSeparator={' '}
                            suffix={' ₽'}
                            decimalSeparator={','}
                            decimalScale={2}
                            fixedDecimalScale={true}
                            renderText={value => <>{value}</>}
                        />
                    </text>
                </svg>
            </div>
            
            <div className={styles.AdsLink}><a href="#">12 объявлений</a></div>
        </div>
    )
}

const Cursor = (prop: CursorProp) => {
    const inverseCursorDirectionDelta: number = 420;
    const cursorClasses = (): string => {
        if (prop.containerWidth - prop.leftOffset > inverseCursorDirectionDelta) {
            return `${styles.Cursor}`;
        } else {
            return [styles.Cursor, styles.CursorRight].join(' ');
        }
    };

    return (
        <div className={cursorClasses()} style={{ left: `${prop.leftOffset}px` }}>
            <div className={styles.Inform}>
                <div className={styles.CurrentAmount}>
                    72 500,00 ₽ — <span className={styles.Percent}>{prop.percent}%</span>
                </div>
                <div className={styles.ToTheGoal}>
                    До цели: <strong>13 месяцев (или 395 дней)</strong><br/>
                    Осталось собрать 2 822 500 ₽
                </div>
                <div className={styles.CurrentResult}>
                    За твои деньги есть уже<br />
                    <a href='#'>3 объявления</a>
                </div>
            </div>
            <img src={CursorIcon} className={styles.Icon}/>
        </div>
    )
}

export const Timeline = ({ }) => {
    const [goalAmount, updateGoalAmount] = useState<number>(2895500.5);
    const [percent, updatePercent] = useState<number>(1.5);


    const timelineRef = useRef<HTMLInputElement | null>(null);
    const [timelinePosition, updateTimelinePosition] = useState<number>(0);

    const calculateProgressWidth = (percent: number, includeOffset: boolean = true): number => {
        const container: DOMRect = timelineRef?.current?.getBoundingClientRect() || {} as DOMRect;
        const shift: number = container.width / 100 * percent || 0;
        return includeOffset ? container.left + shift : shift;
    }

    useEffect(() => {
        updateTimelinePosition(calculateProgressWidth(percent));
    }, [])

    useLayoutEffect(() => {
        const updateTimelineSize = () => {
            updateTimelinePosition(calculateProgressWidth(percent));
        }
        window.addEventListener('resize', updateTimelineSize);
        return () => window.removeEventListener('resize', updateTimelineSize);
    }, [])

    return (
        <div className={styles.Timeline} ref={timelineRef}>
            <Cursor
                percent={percent}
                containerWidth={timelineRef?.current?.getBoundingClientRect().width || 0}
                leftOffset={calculateProgressWidth(percent, false)}
            />

            <div className={styles.TimelineStartFinishContainer}>
                <StartPoint />
                <FinishPoint
                    goalAmount={goalAmount}
                    containerWidth={timelineRef?.current?.getBoundingClientRect().width || 0}
                    leftOffset={calculateProgressWidth(percent, false)}
                />
            </div>

            <div className={styles.TimelineLine}>
                <div style={{width: timelinePosition}} className={styles.TimelineLineProgress}></div>
            </div>
        </div>
    )
}