using UnityEngine;
using System.Collections;
using Util;
using UnityEngine.UI;

public class LoginView : SingletonMonoBehaviour<LoginView>
{
    [Header("玩家名字")]
    public InputField PlayerName;
    [Header("玩家ID")]
    public InputField PlayerId;
    [Header("登陆")]
    public GameObject Login;

    void Start()
    {
        UIEventListener.Get(Login).onClick = (go) =>
        {
            
        };
    }

}
