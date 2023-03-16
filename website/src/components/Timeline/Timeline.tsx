import StartPointImage from '../../assets/StartPoint.png';
import FinishPointImage from '../../assets/FinishPoint.png';
import CursorIcon from '../../assets/CursorIcon.svg';
import styles from './Timeline.module.css'

import { useRecoilValue } from 'recoil';
import { useState, useRef, useEffect, useLayoutEffect } from 'react';
import CurrencyFormat from 'react-currency-format';
import { currentGoalAtom, goalInfoQuery, goalsListQuery } from '../../recoil/goal';



interface CursorProp {
    link: string,
    adsByCurrentAmount: number,
    goalAmount: number,
    accumulatedAmount: number,
    daysUntilBang: number,
    percent: number,
    containerWidth: number,
    leftOffset: number
}

interface FinishPoint {
    goalAmount: number,
    containerWidth: number,
    leftOffset: number
}

const FormatCurrency = (prop: { num: number }) => {
    return (
        <CurrencyFormat
            value={prop.num}
            displayType={'text'}
            thousandSeparator={' '}
            suffix={' ₽'}
            decimalSeparator={','}
            decimalScale={2}
            fixedDecimalScale={true}
            renderText={value => <>{value}</>}
        />
    )
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
                {prop.goalAmount ? 
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
                        <FormatCurrency num={prop.goalAmount} />
                    </text>
                </svg>
                : <></>}
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
                    <FormatCurrency num={prop.accumulatedAmount} /> — <span className={styles.Percent}>{prop.percent}%</span>
                </div>
                <div className={styles.ToTheGoal}>
                    До цели: <strong>{Math.round(prop.daysUntilBang / 30)} месяцев
                    (или {prop.daysUntilBang} дней)</strong><br/>
                    Осталось собрать <FormatCurrency num={(prop.goalAmount - prop.accumulatedAmount) || 0} />
                </div>
                {prop.adsByCurrentAmount ? <div className={styles.CurrentResult}>
                    За твои деньги есть уже<br />
                    <a href={prop.link}>{prop.adsByCurrentAmount} объявления</a>
                </div> : <></>}
            </div>
            <img src={CursorIcon} className={styles.Icon}/>
        </div>
    )
}

export const Timeline = ({ }) => {
    const currentGoal = useRecoilValue(currentGoalAtom);
    const [percent, setPercent] = useState<number>(0);

    const timelineRef = useRef<HTMLInputElement | null>(null);
    const [timelinePosition, updateTimelinePosition] = useState<number>(0);

    const calculateProgressWidth = (percent: number, includeOffset: boolean = true): number => {
        const container: DOMRect = timelineRef?.current?.getBoundingClientRect() || {} as DOMRect;
        const shift: number = container.width / 100 * percent || 0;
        return includeOffset ? container.left + shift : shift;
    }

    useEffect(() => {
        updateTimelinePosition(calculateProgressWidth(percent));
    }, [percent])

    useEffect(() => {
        const persent = Math.round(
            (100 / currentGoal.goal_amount) * (currentGoal.accumulated_amount || 1) * 100
        ) / 100;
        setPercent(persent)
    }, [currentGoal]);

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
                percent={percent || 0}
                containerWidth={timelineRef?.current?.getBoundingClientRect().width || 0}
                leftOffset={calculateProgressWidth(percent, false)}
                accumulatedAmount={currentGoal.accumulated_amount || 0}
                daysUntilBang={currentGoal.days_until_bang || 0}
                adsByCurrentAmount={currentGoal.ads_by_amount || 0}
                link={currentGoal.catalog_url || ''}
                goalAmount={currentGoal.goal_amount}
            />

            <div className={styles.TimelineStartFinishContainer}>
                <StartPoint />
                <FinishPoint
                    goalAmount={currentGoal.goal_amount}
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