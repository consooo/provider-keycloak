#!/bin/bash
# Original Script: https://raw.githubusercontent.com/crossplane/uptest/refs/heads/main/hack/patch-ns.sh
# Added: remove annotation crossplane.io/external-name to Incorrect
function patch {
    kindgroup=$1;
    name=$2;
    namespace=$3;

    # Get all conditions and filter to keep only Test condition
    conditions=$(${KUBECTL} get --namespace "$namespace" "$kindgroup/$name" -o jsonpath='{.status.conditions}' 2>/dev/null || echo "[]")

    # Use jq to filter and keep only Test condition, or use empty array if none exists
    if command -v jq &> /dev/null; then
        test_conditions=$(echo "$conditions" | jq '[.[] | select(.type == "Test")]')
    else
        # Fallback without jq - get Test condition directly
        test_condition=$(${KUBECTL} get --namespace "$namespace" "$kindgroup/$name" -o jsonpath='{.status.conditions[?(@.type=="Test")]}' 2>/dev/null || echo "")
        if [[ -n "$test_condition" ]]; then
            test_conditions="[$test_condition]"
        else
            test_conditions="[]"
        fi
    fi

    # Clear all conditions except Test in a single atomic operation
    if ${KUBECTL} --subresource=status patch --namespace "$namespace" "$kindgroup/$name" --type=merge -p "{\"status\":{\"conditions\":$test_conditions}}" ; then
        return 0;
    else
        return 1;
    fi;
};


kindgroup=$1;
name=$2;
namespace=$3;
attempt=1;
max_attempts=10;
while [[ $attempt -le $max_attempts ]]; do
    if patch "$kindgroup" "$name" "$namespace"; then
        echo "Successfully patched $kindgroup/$name";
        ${KUBECTL} annotate --namespace "$namespace" "$kindgroup/$name" uptest-old-id=$(${KUBECTL} get --namespace "$namespace" "$kindgroup/$name" -o=jsonpath='{.status.atProvider.id}') --overwrite;
        ${KUBECTL} annotate --namespace "$namespace" "$kindgroup/$name" crossplane.io/external-name- --overwrite;
        break;
    else
        printf "Retrying... (%d/%d) for %s/%s/%s\n" "$attempt" "$max_attempts" "$kindgroup" "$name" "$namespace" >&2;
    fi;
    ((attempt++));
    sleep 5;
done;
if [[ $attempt -gt $max_attempts ]]; then
    echo "Failed to patch $kindgroup/$name after $max_attempts attempts";
    exit 1;
fi;
exit 0;