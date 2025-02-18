# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'CustomRuleListArgs',
    'CustomRuleArgs',
    'FrontDoorManagedRuleGroupOverrideArgs',
    'FrontDoorManagedRuleOverrideArgs',
    'FrontDoorManagedRuleSetArgs',
    'FrontDoorMatchConditionArgs',
    'FrontDoorPolicySettingsArgs',
    'ManagedRuleSetListArgs',
]

@pulumi.input_type
class CustomRuleListArgs:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['CustomRuleArgs']]]] = None):
        """
        Defines contents of custom rules
        :param pulumi.Input[Sequence[pulumi.Input['CustomRuleArgs']]] rules: List of rules
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomRuleArgs']]]]:
        """
        List of rules
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class CustomRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[Union[str, 'ActionType']],
                 match_conditions: pulumi.Input[Sequence[pulumi.Input['FrontDoorMatchConditionArgs']]],
                 priority: pulumi.Input[int],
                 rule_type: pulumi.Input[Union[str, 'RuleType']],
                 enabled_state: Optional[pulumi.Input[Union[str, 'CustomRuleEnabledState']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_limit_duration_in_minutes: Optional[pulumi.Input[int]] = None,
                 rate_limit_threshold: Optional[pulumi.Input[int]] = None):
        """
        Defines contents of a web application rule
        :param pulumi.Input[Union[str, 'ActionType']] action: Describes what action to be applied when rule matches.
        :param pulumi.Input[Sequence[pulumi.Input['FrontDoorMatchConditionArgs']]] match_conditions: List of match conditions.
        :param pulumi.Input[int] priority: Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
        :param pulumi.Input[Union[str, 'RuleType']] rule_type: Describes type of rule.
        :param pulumi.Input[Union[str, 'CustomRuleEnabledState']] enabled_state: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        :param pulumi.Input[str] name: Describes the name of the rule.
        :param pulumi.Input[int] rate_limit_duration_in_minutes: Time window for resetting the rate limit count. Default is 1 minute.
        :param pulumi.Input[int] rate_limit_threshold: Number of allowed requests per client within the time window.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "match_conditions", match_conditions)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "rule_type", rule_type)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate_limit_duration_in_minutes is not None:
            pulumi.set(__self__, "rate_limit_duration_in_minutes", rate_limit_duration_in_minutes)
        if rate_limit_threshold is not None:
            pulumi.set(__self__, "rate_limit_threshold", rate_limit_threshold)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[Union[str, 'ActionType']]:
        """
        Describes what action to be applied when rule matches.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[Union[str, 'ActionType']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="matchConditions")
    def match_conditions(self) -> pulumi.Input[Sequence[pulumi.Input['FrontDoorMatchConditionArgs']]]:
        """
        List of match conditions.
        """
        return pulumi.get(self, "match_conditions")

    @match_conditions.setter
    def match_conditions(self, value: pulumi.Input[Sequence[pulumi.Input['FrontDoorMatchConditionArgs']]]):
        pulumi.set(self, "match_conditions", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[Union[str, 'RuleType']]:
        """
        Describes type of rule.
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[Union[str, 'RuleType']]):
        pulumi.set(self, "rule_type", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[Union[str, 'CustomRuleEnabledState']]]:
        """
        Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[Union[str, 'CustomRuleEnabledState']]]):
        pulumi.set(self, "enabled_state", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Describes the name of the rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rateLimitDurationInMinutes")
    def rate_limit_duration_in_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Time window for resetting the rate limit count. Default is 1 minute.
        """
        return pulumi.get(self, "rate_limit_duration_in_minutes")

    @rate_limit_duration_in_minutes.setter
    def rate_limit_duration_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate_limit_duration_in_minutes", value)

    @property
    @pulumi.getter(name="rateLimitThreshold")
    def rate_limit_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Number of allowed requests per client within the time window.
        """
        return pulumi.get(self, "rate_limit_threshold")

    @rate_limit_threshold.setter
    def rate_limit_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate_limit_threshold", value)


@pulumi.input_type
class FrontDoorManagedRuleGroupOverrideArgs:
    def __init__(__self__, *,
                 rule_group_name: pulumi.Input[str],
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleOverrideArgs']]]] = None):
        """
        Defines a managed rule group override setting.
        :param pulumi.Input[str] rule_group_name: Describes the managed rule group to override.
        :param pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleOverrideArgs']]] rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
        """
        pulumi.set(__self__, "rule_group_name", rule_group_name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="ruleGroupName")
    def rule_group_name(self) -> pulumi.Input[str]:
        """
        Describes the managed rule group to override.
        """
        return pulumi.get(self, "rule_group_name")

    @rule_group_name.setter
    def rule_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_group_name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleOverrideArgs']]]]:
        """
        List of rules that will be disabled. If none specified, all rules in the group will be disabled.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleOverrideArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class FrontDoorManagedRuleOverrideArgs:
    def __init__(__self__, *,
                 rule_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[Union[str, 'ActionType']]] = None,
                 enabled_state: Optional[pulumi.Input[Union[str, 'ManagedRuleEnabledState']]] = None):
        """
        Defines a managed rule group override setting.
        :param pulumi.Input[str] rule_id: Identifier for the managed rule.
        :param pulumi.Input[Union[str, 'ActionType']] action: Describes the override action to be applied when rule matches.
        :param pulumi.Input[Union[str, 'ManagedRuleEnabledState']] enabled_state: Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
        """
        pulumi.set(__self__, "rule_id", rule_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        """
        Identifier for the managed rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Union[str, 'ActionType']]]:
        """
        Describes the override action to be applied when rule matches.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Union[str, 'ActionType']]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[Union[str, 'ManagedRuleEnabledState']]]:
        """
        Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[Union[str, 'ManagedRuleEnabledState']]]):
        pulumi.set(self, "enabled_state", value)


@pulumi.input_type
class FrontDoorManagedRuleSetArgs:
    def __init__(__self__, *,
                 rule_set_type: pulumi.Input[str],
                 rule_set_version: pulumi.Input[str],
                 rule_group_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleGroupOverrideArgs']]]] = None):
        """
        Defines a managed rule set.
        :param pulumi.Input[str] rule_set_type: Defines the rule set type to use.
        :param pulumi.Input[str] rule_set_version: Defines the version of the rule set to use.
        :param pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleGroupOverrideArgs']]] rule_group_overrides: Defines the rule group overrides to apply to the rule set.
        """
        pulumi.set(__self__, "rule_set_type", rule_set_type)
        pulumi.set(__self__, "rule_set_version", rule_set_version)
        if rule_group_overrides is not None:
            pulumi.set(__self__, "rule_group_overrides", rule_group_overrides)

    @property
    @pulumi.getter(name="ruleSetType")
    def rule_set_type(self) -> pulumi.Input[str]:
        """
        Defines the rule set type to use.
        """
        return pulumi.get(self, "rule_set_type")

    @rule_set_type.setter
    def rule_set_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_set_type", value)

    @property
    @pulumi.getter(name="ruleSetVersion")
    def rule_set_version(self) -> pulumi.Input[str]:
        """
        Defines the version of the rule set to use.
        """
        return pulumi.get(self, "rule_set_version")

    @rule_set_version.setter
    def rule_set_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_set_version", value)

    @property
    @pulumi.getter(name="ruleGroupOverrides")
    def rule_group_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleGroupOverrideArgs']]]]:
        """
        Defines the rule group overrides to apply to the rule set.
        """
        return pulumi.get(self, "rule_group_overrides")

    @rule_group_overrides.setter
    def rule_group_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleGroupOverrideArgs']]]]):
        pulumi.set(self, "rule_group_overrides", value)


@pulumi.input_type
class FrontDoorMatchConditionArgs:
    def __init__(__self__, *,
                 match_value: pulumi.Input[Sequence[pulumi.Input[str]]],
                 match_variable: pulumi.Input[Union[str, 'FrontDoorMatchVariable']],
                 operator: pulumi.Input[Union[str, 'Operator']],
                 negate_condition: Optional[pulumi.Input[bool]] = None,
                 selector: Optional[pulumi.Input[str]] = None,
                 transforms: Optional[pulumi.Input[Sequence[pulumi.Input[Union[str, 'TransformType']]]]] = None):
        """
        Define a match condition.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] match_value: List of possible match values.
        :param pulumi.Input[Union[str, 'FrontDoorMatchVariable']] match_variable: Request variable to compare with.
        :param pulumi.Input[Union[str, 'Operator']] operator: Comparison type to use for matching with the variable value.
        :param pulumi.Input[bool] negate_condition: Describes if the result of this condition should be negated.
        :param pulumi.Input[str] selector: Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is null.
        :param pulumi.Input[Sequence[pulumi.Input[Union[str, 'TransformType']]]] transforms: List of transforms.
        """
        pulumi.set(__self__, "match_value", match_value)
        pulumi.set(__self__, "match_variable", match_variable)
        pulumi.set(__self__, "operator", operator)
        if negate_condition is not None:
            pulumi.set(__self__, "negate_condition", negate_condition)
        if selector is not None:
            pulumi.set(__self__, "selector", selector)
        if transforms is not None:
            pulumi.set(__self__, "transforms", transforms)

    @property
    @pulumi.getter(name="matchValue")
    def match_value(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of possible match values.
        """
        return pulumi.get(self, "match_value")

    @match_value.setter
    def match_value(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "match_value", value)

    @property
    @pulumi.getter(name="matchVariable")
    def match_variable(self) -> pulumi.Input[Union[str, 'FrontDoorMatchVariable']]:
        """
        Request variable to compare with.
        """
        return pulumi.get(self, "match_variable")

    @match_variable.setter
    def match_variable(self, value: pulumi.Input[Union[str, 'FrontDoorMatchVariable']]):
        pulumi.set(self, "match_variable", value)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[Union[str, 'Operator']]:
        """
        Comparison type to use for matching with the variable value.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[Union[str, 'Operator']]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter(name="negateCondition")
    def negate_condition(self) -> Optional[pulumi.Input[bool]]:
        """
        Describes if the result of this condition should be negated.
        """
        return pulumi.get(self, "negate_condition")

    @negate_condition.setter
    def negate_condition(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "negate_condition", value)

    @property
    @pulumi.getter
    def selector(self) -> Optional[pulumi.Input[str]]:
        """
        Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is null.
        """
        return pulumi.get(self, "selector")

    @selector.setter
    def selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "selector", value)

    @property
    @pulumi.getter
    def transforms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union[str, 'TransformType']]]]]:
        """
        List of transforms.
        """
        return pulumi.get(self, "transforms")

    @transforms.setter
    def transforms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union[str, 'TransformType']]]]]):
        pulumi.set(self, "transforms", value)


@pulumi.input_type
class FrontDoorPolicySettingsArgs:
    def __init__(__self__, *,
                 custom_block_response_body: Optional[pulumi.Input[str]] = None,
                 custom_block_response_status_code: Optional[pulumi.Input[int]] = None,
                 enabled_state: Optional[pulumi.Input[Union[str, 'PolicyEnabledState']]] = None,
                 mode: Optional[pulumi.Input[Union[str, 'PolicyMode']]] = None,
                 redirect_url: Optional[pulumi.Input[str]] = None):
        """
        Defines top-level WebApplicationFirewallPolicy configuration settings.
        :param pulumi.Input[str] custom_block_response_body: If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
        :param pulumi.Input[int] custom_block_response_status_code: If the action type is block, customer can override the response status code.
        :param pulumi.Input[Union[str, 'PolicyEnabledState']] enabled_state: Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
        :param pulumi.Input[Union[str, 'PolicyMode']] mode: Describes if it is in detection mode or prevention mode at policy level.
        :param pulumi.Input[str] redirect_url: If action type is redirect, this field represents redirect URL for the client.
        """
        if custom_block_response_body is not None:
            pulumi.set(__self__, "custom_block_response_body", custom_block_response_body)
        if custom_block_response_status_code is not None:
            pulumi.set(__self__, "custom_block_response_status_code", custom_block_response_status_code)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if redirect_url is not None:
            pulumi.set(__self__, "redirect_url", redirect_url)

    @property
    @pulumi.getter(name="customBlockResponseBody")
    def custom_block_response_body(self) -> Optional[pulumi.Input[str]]:
        """
        If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
        """
        return pulumi.get(self, "custom_block_response_body")

    @custom_block_response_body.setter
    def custom_block_response_body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_block_response_body", value)

    @property
    @pulumi.getter(name="customBlockResponseStatusCode")
    def custom_block_response_status_code(self) -> Optional[pulumi.Input[int]]:
        """
        If the action type is block, customer can override the response status code.
        """
        return pulumi.get(self, "custom_block_response_status_code")

    @custom_block_response_status_code.setter
    def custom_block_response_status_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "custom_block_response_status_code", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[Union[str, 'PolicyEnabledState']]]:
        """
        Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[Union[str, 'PolicyEnabledState']]]):
        pulumi.set(self, "enabled_state", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[Union[str, 'PolicyMode']]]:
        """
        Describes if it is in detection mode or prevention mode at policy level.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[Union[str, 'PolicyMode']]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="redirectUrl")
    def redirect_url(self) -> Optional[pulumi.Input[str]]:
        """
        If action type is redirect, this field represents redirect URL for the client.
        """
        return pulumi.get(self, "redirect_url")

    @redirect_url.setter
    def redirect_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_url", value)


@pulumi.input_type
class ManagedRuleSetListArgs:
    def __init__(__self__, *,
                 managed_rule_sets: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleSetArgs']]]] = None):
        """
        Defines the list of managed rule sets for the policy.
        :param pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleSetArgs']]] managed_rule_sets: List of rule sets.
        """
        if managed_rule_sets is not None:
            pulumi.set(__self__, "managed_rule_sets", managed_rule_sets)

    @property
    @pulumi.getter(name="managedRuleSets")
    def managed_rule_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleSetArgs']]]]:
        """
        List of rule sets.
        """
        return pulumi.get(self, "managed_rule_sets")

    @managed_rule_sets.setter
    def managed_rule_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FrontDoorManagedRuleSetArgs']]]]):
        pulumi.set(self, "managed_rule_sets", value)


