using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Util;
public class EventCenterManager : Singleton<EventCenterManager>
{
    public delegate void DelegateActon(params object[] objs);

    Dictionary<int, List<DelegateActon>> EventMap = new Dictionary<int, List<DelegateActon>>();

    public void Send(GlobalEventEnum globalEventEnum,params object[] objs)
    {
        Send((int)globalEventEnum, objs);
    }
    public void Send(M2CEventEnum m2cEventEnum, params object[] objs)
    {
        Send((int)m2cEventEnum, objs);
    }
    /// <summary>
    /// 发送消息
    /// </summary>
    /// <param name="eventEnum"></param>
    /// <param name="objs"></param>
    public void Send(int eventEnum, params object[] objs)
    {
        List<DelegateActon> actionGroup = null;
        if (EventMap.TryGetValue(eventEnum, out actionGroup))
        {
            for (int i = actionGroup.Count - 1; i >= 0; i--)
            {
                actionGroup[i](objs);
            }
        }
    }
    public void RegisterEvent(GlobalEventEnum eventEnum, DelegateActon action)
    {
        RegisterEvent((int)eventEnum, action);
    }
    public void RegisterEvent(M2CEventEnum eventEnum, DelegateActon action)
    {
        RegisterEvent((int)eventEnum, action);
    }
    /// <summary>
    /// 注册消息
    /// </summary>
    /// <param name="eventEnum"></param>
    /// <param name="action"></param>
    public void RegisterEvent(int eventEnum, DelegateActon action)
    {
        List<DelegateActon> actionGroup = null;
        if (!EventMap.TryGetValue(eventEnum, out actionGroup))
        {
            actionGroup = new List<DelegateActon>();
            EventMap.Add((int)eventEnum, actionGroup);
        }

        if (!actionGroup.Contains(action))
        {
            actionGroup.Add(action);
        }

    }
    public void UnRegisterEvent(GlobalEventEnum eventEnum, DelegateActon action)
    {
        UnRegisterEvent((int)eventEnum, action);
    }
    public void UnRegisterEvent(M2CEventEnum eventEnum, DelegateActon action)
    {
        UnRegisterEvent((int)eventEnum, action);
    }
    /// <summary>
    /// 注销事件
    /// </summary>
    /// <param name="eventEnum"></param>
    /// <param name="action"></param>
    public void UnRegisterEvent(int eventEnum, DelegateActon action)
    {
        List<DelegateActon> actionGroup = null;
        if (EventMap.TryGetValue(eventEnum, out actionGroup))
        {
            for (int i = actionGroup.Count - 1; i > 0; i--)
            {
                if (actionGroup[i] == action)
                    actionGroup.RemoveAt(i);
            }
        }
    }
}

/// <summary>
/// 事件类型
/// </summary>
public enum GlobalEventEnum
{
    None,
}

public enum M2CEventEnum
{
    None = 10000,
}