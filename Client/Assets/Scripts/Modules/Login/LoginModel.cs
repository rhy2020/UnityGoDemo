using UnityEngine;
using System.Collections;
using Msg;

public class LoginModel : BaseModel<LoginModel>
{
    protected override void InitAddTocHandler()
    {
        AddTocHandler(typeof(LoginSuccessfull), LoginSuccessfull);
        AddTocHandler(typeof(LoginFaild), LoginFaild);

    }
    private void LoginSuccessfull(object data)
    {
        LoginSuccessfull loginSuccessfull = data as LoginSuccessfull;
        
    }
    private void LoginFaild(object data)
    {
        LoginFaild loginFaild = data as LoginFaild;

    }
}
