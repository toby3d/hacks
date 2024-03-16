package http

// 700 HTTP status code series as known unknowns.
// See: https://github.com/joho/7XX-rfc
const (
	// Inexcusable
	StatusMeh                                                       = 701
	StatusEmacs                                                     = 702
	StatusExplosion                                                 = 703
	StatusGotoFail                                                  = 704
	StatusIWroteTheCodeAndMissedTheNecessaryValidationByAnOversight = 705 // See [StatusHeartbleed]
	StatusDeleteYourAccount                                         = 706
	StatusCantQuitVi                                                = 707

	// Novelty Implementations
	StatusPHP              = 710
	StatusConvenienceStore = 711
	StatusNoSQL            = 712
	StatusIAmNotATeapot    = 718
	StatusHaskell          = 719

	// Edge Cases
	StatusUnpossible                  = 720
	StatusKnownUnknowns               = 721
	StatusUnknownUnknowns             = 722
	StatusTricky                      = 723
	StatusThisLineShouldBeUnreachable = 724
	StatusItWorksOnMyMachine          = 725
	StatusItsAFeatureNotABug          = 726
	Status32BitsIsPlenty              = 727
	StatusItWorksInMyTimezone         = 728

	// Fucking
	StatusFuckingNpm                 = 730
	StatusFuckingRubygems            = 731
	StatusFuckingUnicode             = 732
	StatusFuckingDeadlocks           = 733
	StatusFuckingDeferreds           = 734
	StatusFuckingRaceConditions      = 736
	StatusFuckingIE                  = 735
	StatusFuckThreadsing             = 737
	StatusFuckingExactlyOnceDelivery = 738
	StatusFuckingWindows             = 739

	// Reserved for meritocracy related bullshit
	// 74x TBD.  Got the brains trust on the case.

	// Syntax Errors
	StatusDidntBotherToCompileIt           = 750
	StatusSyntaxError                      = 753
	StatusTooManySemiColons                = 754
	StatusNotEnoughSemiColons              = 755
	StatusInsufficientlyPolite             = 756
	StatusExcessivelyPolite                = 757
	StatusUnexpectedT_PAAMAYIM_NEKUDOTAYIM = 759

	// Substance-Affected Developer
	StatusHungover                                                            = 761
	StatusStoned                                                              = 762
	StatusUnderCaffeinated                                                    = 763
	StatusOverCaffeinated                                                     = 764
	StatusRailscamp                                                           = 765
	StatusSober                                                               = 766
	StatusDrunk                                                               = 767
	StatusAccidentallyTookSleepingPillsInsteadOfMigrainePillsDuringCrunchWeek = 768

	// Predictable Problems
	StatusCachedForTooLong         = 771
	StatusNotCachedLongEnough      = 772
	StatusNotCachedAtAll           = 773
	StatusWhyWasThisCached         = 774
	StatusOutOfCash                = 775
	StatusErrorOnTheException      = 776
	StatusCoincidence              = 777
	StatusOffByOneError            = 778
	StatusOffByTooManyToCountError = 779

	// Somebody Else's Problem
	StatusProjectOwnerNotResponding     = 780
	StatusOperations                    = 781
	StatusQA                            = 782
	StatusItWasACustomerRequestHonestly = 783
	StatusManagementObviously           = 784
	StatusTPSCoverSheetNotAttached      = 785
	StatusTryItNow                      = 786
	StatusFurtherFundingRequired        = 787
	StatusDesignersFinalDesignsWerent   = 788
	StatusNotMyDepartment               = 789

	// Internet crashed
	StatusTheInternetShutDownDueToCopyrightRestrictions     = 791
	StatusClimateChangeDrivenCatastrophicWeatherEvent       = 792
	StatusZombieApocalypse                                  = 793
	StatusSomeoneLetPGNearAREPL                             = 794
	StatusHeartbleed                                        = 795 // See [StatusIWroteTheCodeAndMissedTheNecessaryValidationByAnOversight]
	StatusSomeDNSFuckeryIdno                                = 796
	StatusThisIsTheLastPageOfTheInternetGoBack              = 797
	StatusICheckedTheDbBackupsCupboardAndTheCupboardWasBare = 798
	StatusEndOfTheWorld                                     = 799
)
