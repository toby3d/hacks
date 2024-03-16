// Package http contains utilities, jokes and helpers for manipulate HTTP data
// in clients and servers.
package http

import "net/http"

// StatusText wraps native [net/http.StatusText], but considers the status codes
// provided in this package before referring to the standard ones.
func StatusText(code int) string {
	switch code {
	default:
		// Reserved for meritocracy related bullshit
		if StatusFuckingWindows < code && code < StatusDidntBotherToCompileIt {
			return "TBD. Got the brains trust on the case."
		}

		return http.StatusText(code)
	case StatusQuestionableButOkay:
		return "Questionable But Okay"
	case StatusMeh:
		return "Meh"
	case StatusEmacs:
		return "Emacs"
	case StatusExplosion:
		return "Explosion"
	case StatusGotoFail:
		return "Goto Fail"
	case StatusIWroteTheCodeAndMissedTheNecessaryValidationByAnOversight:
		return "I wrote the code and missed the necessary validation by an oversight"
	case StatusDeleteYourAccount:
		return "Delete Your Account"
	case StatusCantQuitVi:
		return "Can't quit vi"
	case StatusPHP:
		return "PHP"
	case StatusConvenienceStore:
		return "Convenience Store"
	case StatusNoSQL:
		return "NoSQL"
	case StatusIAmNotATeapot:
		return "I am not a teapot"
	case StatusHaskell:
		return "Haskell"
	case StatusUnpossible:
		return "Unpossible"
	case StatusKnownUnknowns:
		return "Known Unknowns"
	case StatusUnknownUnknowns:
		return "Unknown Unknowns"
	case StatusTricky:
		return "Tricky"
	case StatusThisLineShouldBeUnreachable:
		return "This line should be unreachable"
	case StatusItWorksOnMyMachine:
		return "It works on my machine"
	case StatusItsAFeatureNotABug:
		return "It's a feature, not a bug"
	case Status32BitsIsPlenty:
		return "32 bits is plenty"
	case StatusItWorksInMyTimezone:
		return "It works in my timezone"
	case StatusFuckingNpm:
		return "Fucking npm"
	case StatusFuckingRubygems:
		return "Fucking Rubygems"
	case StatusFuckingUnicode:
		return "Fucking Unic&#128169;de"
	case StatusFuckingDeadlocks:
		return "Fucking Deadlocks"
	case StatusFuckingDeferreds:
		return "Fucking Deferreds"
	case StatusFuckingRaceConditions:
		return "Fucking Race Conditions"
	case StatusFuckingIE:
		return "Fucking IE"
	case StatusFuckThreadsing:
		return "FuckThreadsing"
	case StatusFuckingExactlyOnceDelivery:
		return "Fucking Exactly-once Delivery"
	case StatusFuckingWindows:
		return "Fucking Windows"
	case StatusDidntBotherToCompileIt:
		return "Didn't bother to compile it"
	case StatusSyntaxError:
		return "Syntax Error"
	case StatusTooManySemiColons:
		return "Too many semi-colons"
	case StatusNotEnoughSemiColons:
		return "Not enough semi-colons"
	case StatusInsufficientlyPolite:
		return "Insufficiently polite"
	case StatusExcessivelyPolite:
		return "Excessively polite"
	case StatusUnexpectedT_PAAMAYIM_NEKUDOTAYIM:
		return "Unexpected \"T_PAAMAYIM_NEKUDOTAYIM\""
	case StatusHungover:
		return "Hungover"
	case StatusStoned:
		return "Stoned"
	case StatusUnderCaffeinated:
		return "Under-Caffeinated"
	case StatusOverCaffeinated:
		return "Over-Caffeinated"
	case StatusRailscamp:
		return "Railscamp"
	case StatusSober:
		return "Sober"
	case StatusDrunk:
		return "Drunk"
	case StatusAccidentallyTookSleepingPillsInsteadOfMigrainePillsDuringCrunchWeek:
		return "Accidentally Took Sleeping Pills Instead Of Migraine Pills During Crunch Week"
	case StatusCachedForTooLong:
		return "Cached for too long"
	case StatusNotCachedLongEnough:
		return "Not cached long enough"
	case StatusNotCachedAtAll:
		return "Not cached at all"
	case StatusWhyWasThisCached:
		return "Why was this cached?"
	case StatusOutOfCash:
		return "Out of cash"
	case StatusErrorOnTheException:
		return "Error on the Exception"
	case StatusCoincidence:
		return "Coincidence"
	case StatusOffByOneError:
		return "Off By One Error"
	case StatusOffByTooManyToCountError:
		return "Off By Too Many To Count Error"
	case StatusProjectOwnerNotResponding:
		return "Project owner not responding"
	case StatusOperations:
		return "Operations"
	case StatusQA:
		return "QA"
	case StatusItWasACustomerRequestHonestly:
		return "It was a customer request, honestly"
	case StatusManagementObviously:
		return "Management, obviously"
	case StatusTPSCoverSheetNotAttached:
		return "TPS Cover Sheet not attached"
	case StatusTryItNow:
		return "Try it now"
	case StatusFurtherFundingRequired:
		return "Further Funding Required"
	case StatusDesignersFinalDesignsWerent:
		return "Designer's final designs weren't"
	case StatusNotMyDepartment:
		return "Not my department"
	case StatusTheInternetShutDownDueToCopyrightRestrictions:
		return "The Internet shut down due to copyright restrictions"
	case StatusClimateChangeDrivenCatastrophicWeatherEvent:
		return "Climate change driven catastrophic weather event"
	case StatusZombieApocalypse:
		return "Zombie Apocalypse"
	case StatusSomeoneLetPGNearAREPL:
		return "Someone let PG near a REPL"
	case StatusHeartbleed:
		return "#heartbleed"
	case StatusSomeDNSFuckeryIdno:
		return "Some DNS fuckery idno"
	case StatusThisIsTheLastPageOfTheInternetGoBack:
		return "This is the last page of the Internet.  Go back"
	case StatusICheckedTheDbBackupsCupboardAndTheCupboardWasBare:
		return "I checked the db backups cupboard and the cupboard was bare"
	case StatusEndOfTheWorld:
		return "End of the world"
	}
}
