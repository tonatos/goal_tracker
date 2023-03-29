import { useRecoilState, useSetRecoilState } from "recoil";
import { useFetchWrapper } from "../helpers";
import { CreateUpdateGoal } from "../interfaces/goal.interfaces";
import { currentGoalAtom, goalsAtom } from "../recoil/goal";

export const useGoalActions = () => {
    const baseUrl = `${process.env.REACT_APP_API_URL}/goal`;
    const fetchWrapper = useFetchWrapper();
    const [ goals, setGoals ] = useRecoilState(goalsAtom);
    const setCurrentGoal = useSetRecoilState(currentGoalAtom);

    const getAll = async () => {
        const r = await fetchWrapper.get(baseUrl);
        setGoals(r.data)
        return r.data;
    }

    const getDefault =  async (id: number | undefined) => {
        if (!id) {
            return;
        }
        const r = await fetchWrapper.get(`${baseUrl}/${id}`);
        setCurrentGoal(r.data);
        return r.data;
    }

    const create = async (goal: CreateUpdateGoal) => {
        if (!goal) {
            return;
        }
        const r = await fetchWrapper.post(`${baseUrl}/`, goal);
        await getAll();
        setCurrentGoal(r.data);
        return r.data;
    }

    const del = async (id: number | undefined) => {
        if (!id) {
            return;
        }
        const r = await fetchWrapper.delete(`${baseUrl}/${id}/`);
        await getAll();
        return;
    }
    
    return {
        getAll,
        getDefault,
        create,
        del,
    }
}