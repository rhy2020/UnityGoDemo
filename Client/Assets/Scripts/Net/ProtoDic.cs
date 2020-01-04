
using Google.Protobuf;
using Msg;
using System;
using System.Collections.Generic;

namespace Proto
{
   public class ProtoDic
   {
       private static List<int> _protoId = new List<int>
       {
            0,
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8,
            9,
            10,
            11,
            12,
            13,
            14,
            15,
            16,
            17,
            18,
            19,
            20,
            21,
            22,
            23,
            24,
            25,
            26,
            27,
            28,
            29,
            30,
        };

      private static List<Type>_protoType = new List<Type>
      {
            typeof(RspHead),
            typeof(HeartReq),
            typeof(HeartRsp),
            typeof(StartFight),
            typeof(FightResult),
            typeof(EnterFight),
            typeof(SignUpResponse),
            typeof(TosChat),
            typeof(TocChat),
            typeof(Login),
            typeof(PlayerBaseInfo),
            typeof(LoginSuccessfull),
            typeof(LoginFaild),
            typeof(RoomData),
            typeof(RoomListReq),
            typeof(RoomListRsp),
            typeof(JoinRoomReq),
            typeof(JoinRoomRsp),
            typeof(LeaveRoomReq),
            typeof(LeaveRoomRsp),
            typeof(RoomPlayerJoinNtf),
            typeof(RoomPlayerLeaveNtf),
            typeof(RoomBattleReadyReq),
            typeof(RoomBattleReadyRsp),
            typeof(RoomBattleReadyNtf),
            typeof(RoomBattleCancleReadyReq),
            typeof(RoomBattleCancleReadyRsp),
            typeof(RoomBattleCancleReadyNtf),
            typeof(RoomBattleStartReq),
            typeof(RoomBattleStartRsp),
            typeof(RoomBattleStartNtf),
       };

       private static readonly Dictionary<RuntimeTypeHandle, MessageParser> Parsers = new Dictionary<RuntimeTypeHandle, MessageParser>()
       {
            {typeof(RspHead).TypeHandle,RspHead.Parser },
            {typeof(HeartReq).TypeHandle,HeartReq.Parser },
            {typeof(HeartRsp).TypeHandle,HeartRsp.Parser },
            {typeof(StartFight).TypeHandle,StartFight.Parser },
            {typeof(FightResult).TypeHandle,FightResult.Parser },
            {typeof(EnterFight).TypeHandle,EnterFight.Parser },
            {typeof(SignUpResponse).TypeHandle,SignUpResponse.Parser },
            {typeof(TosChat).TypeHandle,TosChat.Parser },
            {typeof(TocChat).TypeHandle,TocChat.Parser },
            {typeof(Login).TypeHandle,Login.Parser },
            {typeof(PlayerBaseInfo).TypeHandle,PlayerBaseInfo.Parser },
            {typeof(LoginSuccessfull).TypeHandle,LoginSuccessfull.Parser },
            {typeof(LoginFaild).TypeHandle,LoginFaild.Parser },
            {typeof(RoomData).TypeHandle,RoomData.Parser },
            {typeof(RoomListReq).TypeHandle,RoomListReq.Parser },
            {typeof(RoomListRsp).TypeHandle,RoomListRsp.Parser },
            {typeof(JoinRoomReq).TypeHandle,JoinRoomReq.Parser },
            {typeof(JoinRoomRsp).TypeHandle,JoinRoomRsp.Parser },
            {typeof(LeaveRoomReq).TypeHandle,LeaveRoomReq.Parser },
            {typeof(LeaveRoomRsp).TypeHandle,LeaveRoomRsp.Parser },
            {typeof(RoomPlayerJoinNtf).TypeHandle,RoomPlayerJoinNtf.Parser },
            {typeof(RoomPlayerLeaveNtf).TypeHandle,RoomPlayerLeaveNtf.Parser },
            {typeof(RoomBattleReadyReq).TypeHandle,RoomBattleReadyReq.Parser },
            {typeof(RoomBattleReadyRsp).TypeHandle,RoomBattleReadyRsp.Parser },
            {typeof(RoomBattleReadyNtf).TypeHandle,RoomBattleReadyNtf.Parser },
            {typeof(RoomBattleCancleReadyReq).TypeHandle,RoomBattleCancleReadyReq.Parser },
            {typeof(RoomBattleCancleReadyRsp).TypeHandle,RoomBattleCancleReadyRsp.Parser },
            {typeof(RoomBattleCancleReadyNtf).TypeHandle,RoomBattleCancleReadyNtf.Parser },
            {typeof(RoomBattleStartReq).TypeHandle,RoomBattleStartReq.Parser },
            {typeof(RoomBattleStartRsp).TypeHandle,RoomBattleStartRsp.Parser },
            {typeof(RoomBattleStartNtf).TypeHandle,RoomBattleStartNtf.Parser },
       };

        public static MessageParser GetMessageParser(RuntimeTypeHandle typeHandle)
        {
            MessageParser messageParser;
            Parsers.TryGetValue(typeHandle, out messageParser);
            return messageParser;
        }

        public static Type GetProtoTypeByProtoId(int protoId)
        {
            int index = _protoId.IndexOf(protoId);
            return _protoType[index];
        }

        public static int GetProtoIdByProtoType(Type type)
        {
            int index = _protoType.IndexOf(type);
            return _protoId[index];
        }

        public static bool ContainProtoId(int protoId)
        {
            if(_protoId.Contains(protoId))
            {
                return true;
            }
            return false;
        }

        public static bool ContainProtoType(Type type)
        {
            if(_protoType.Contains(type))
            {
                return true;
            }
            return false;
        }
    }
}