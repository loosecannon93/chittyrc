/*
 *
 *  chirc
 *
 *  Reply codes
 *
 */
package main

const (
	RPL_WELCOME  = "001"
	RPL_YOURHOST = "002"
	RPL_CREATED  = "003"
	RPL_MYINFO   = "004"

	RPL_LUSERCLIENT   = "251"
	RPL_LUSEROP       = "252"
	RPL_LUSERUNKNOWN  = "253"
	RPL_LUSERCHANNELS = "254"
	RPL_LUSERME       = "255"

	RPL_AWAY    = "301"
	RPL_UNAWAY  = "305"
	RPL_NOWAWAY = "306"

	RPL_WHOISUSER     = "311"
	RPL_WHOISSERVER   = "312"
	RPL_WHOISOPERATOR = "313"
	RPL_WHOISIDLE     = "317"
	RPL_ENDOFWHOIS    = "318"
	RPL_WHOISCHANNELS = "319"

	RPL_WHOREPLY = "352"
	RPL_ENDOFWHO = "315"

	RPL_LIST    = "322"
	RPL_LISTEND = "323"

	RPL_CHANNELMODEIS = "324"

	RPL_NOTOPIC = "331"
	RPL_TOPIC   = "332"

	RPL_NAMREPLY   = "353"
	RPL_ENDOFNAMES = "366"

	RPL_MOTDSTART = "375"
	RPL_MOTD      = "372"
	RPL_ENDOFMOTD = "376"

	RPL_YOUREOPER = "381"

	ERR_NOSUCHNICK       = "401"
	ERR_NOSUCHCHANNEL    = "403"
	ERR_CANNOTSENDTOCHAN = "404"
	ERR_UNKNOWNCOMMAND   = "421"
	ERR_NOMOTD           = "422"
	ERR_NONICKNAMEGIVEN  = "431"
	ERR_NICKNAMEINUSE    = "433"
	ERR_USERNOTINCHANNEL = "441"
	ERR_NOTONCHANNEL     = "442"
	ERR_NOTREGISTERED    = "451"
	ERR_NEEDMOREPARAMS   = "461"
	ERR_ALREADYREGISTRED = "462"
	ERR_PASSWDMISMATCH   = "464"
	ERR_UNKNOWNMODE      = "472"
	ERR_CHANOPRIVSNEEDED = "482"
	ERR_UMODEUNKNOWNFLAG = "501"
	ERR_USERSDONTMATCH   = "502"
)
